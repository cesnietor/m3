// This file is part of MinIO Kubernetes Cloud
// Copyright (c) 2019 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package api

import (
	"context"
	"fmt"
	"log"

	pb "github.com/minio/m3/api/stubs"
	"github.com/minio/m3/cluster"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateServiceAccount Creates a new service account and assigns to it the permissions selected
func (s *server) CreateServiceAccount(ctx context.Context, in *pb.CreateServiceAccountRequest) (res *pb.CreateServiceAccountResponse, err error) {
	name := in.GetName()
	permisionsIDs := in.GetPermissionIds()
	if name == "" {
		return nil, status.New(codes.InvalidArgument, "a name is needed").Err()
	}
	if len(permisionsIDs) == 0 {
		return nil, status.New(codes.InvalidArgument, "a list of permissions is needed").Err()
	}
	// start app context
	appCtx, err := cluster.NewTenantContextWithGrpcContext(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			log.Println(err.Error())
			appCtx.Rollback()
			return
		}
		// if no error happened to this point commit transaction
		err = appCtx.Commit()
	}()

	serviceAccount, saCred, err := cluster.AddServiceAccount(appCtx, appCtx.Tenant.ShortName, name, &name)
	if err != nil {
		log.Println(err.Error())
		return nil, status.New(codes.Internal, "error creating service account").Err()
	}

	var permissionIDsArr []*uuid.UUID
	for _, idString := range permisionsIDs {
		permUUID, err := uuid.FromString(idString)
		if err != nil {
			return nil, status.New(codes.Internal, "permission uuid not valid").Err()
		}
		permissionIDsArr = append(permissionIDsArr, &permUUID)
	}
	// perform actions
	err = cluster.AssignMultiplePermissionsToSA(appCtx, &serviceAccount.ID, permissionIDsArr)
	if err != nil {
		log.Println(err.Error())
		return nil, status.New(codes.Internal, "Internal error").Err()
	}

	return &pb.CreateServiceAccountResponse{
		ServiceAccount: &pb.ServiceAccount{
			Id:        serviceAccount.ID.String(),
			Name:      serviceAccount.Name,
			AccessKey: serviceAccount.AccessKey,
			Enabled:   serviceAccount.Enabled,
		},
		SecretKey: saCred.SecretKey,
	}, nil
}

// ListServiceAccounts lists all service accounts of a tenant
func (s *server) ListServiceAccounts(ctx context.Context, in *pb.ListServiceAccountsRequest) (res *pb.ListServiceAccountsResponse, err error) {
	offset := in.GetOffset()
	limit := in.GetLimit()
	if limit == 0 {
		limit = defaultRequestLimit
	}
	// start app context
	appCtx, err := cluster.NewTenantContextWithGrpcContext(ctx)
	if err != nil {
		return nil, err
	}
	serviceAccounts, err := cluster.GetServiceAccountList(appCtx, int(offset), int(limit))
	if err != nil {
		log.Println(err.Error())
		return nil, status.New(codes.Internal, "Internal error").Err()
	}

	var servAccountsResp []*pb.ServiceAccount
	for _, serviceAccount := range serviceAccounts {
		sa := &pb.ServiceAccount{
			Id:        serviceAccount.ID.String(),
			Name:      serviceAccount.Name,
			AccessKey: serviceAccount.AccessKey,
			Enabled:   serviceAccount.Enabled,
		}
		servAccountsResp = append(servAccountsResp, sa)
	}
	return &pb.ListServiceAccountsResponse{
		ServiceAccounts: servAccountsResp,
		Total:           int32(len(servAccountsResp)),
	}, nil
}

func (s *server) UpdateServiceAccount(ctx context.Context, in *pb.UpdateServiceAccountRequest) (res *pb.InfoServiceAccountResponse, err error) {
	idRequest := in.GetId()
	nameRequest := in.GetName()
	enabledRequest := in.GetEnabled()
	permisionsIDs := in.GetPermissionIds()
	if idRequest == "" {
		return nil, status.New(codes.InvalidArgument, "an id is needed").Err()
	}
	if nameRequest == "" {
		return nil, status.New(codes.InvalidArgument, "an name is needed").Err()
	}
	if len(permisionsIDs) == 0 {
		return nil, status.New(codes.InvalidArgument, "a list of permissions is needed").Err()
	}
	// start app context
	appCtx, err := cluster.NewTenantContextWithGrpcContext(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			appCtx.Rollback()
			return
		}
	}()

	// Fetch the service account
	id, err := uuid.FromString(idRequest)
	if err != nil {
		log.Println(err.Error())
		return nil, status.New(codes.InvalidArgument, "not valid id").Err()
	}
	serviceAccount, err := cluster.GetServiceAccountByID(appCtx, id)
	if err != nil {
		fmt.Println(err.Error())
		log.Println(err.Error())
		return nil, status.New(codes.NotFound, "service account not found").Err()
	}
	// get all the permissions for the service account
	perms, err := cluster.GetAllThePermissionForServiceAccount(appCtx, &serviceAccount.ID)
	if err != nil {
		fmt.Println(err.Error())
		log.Println(err.Error())
		return nil, status.New(codes.Internal, "Internal error").Err()
	}

	// Compare current Permissions with the desired ones

	var currentPerms []string
	for _, perm := range perms {
		currentPerms = append(currentPerms, perm.ID.String())
	}
	fmt.Println("currentPerms: ", currentPerms)
	// TODO: parallelize
	permissionsToCreate := differenceArrays(permisionsIDs, currentPerms)
	permissionsToDelete := differenceArrays(currentPerms, permisionsIDs)

	// Create new service_accounts_permissions
	permsToCreateIDs, err := cluster.UUIDsFromStringArr(permissionsToCreate)
	if err != nil {
		log.Println(err.Error())
		return nil, status.New(codes.InvalidArgument, "invalid permission id").Err()
	}
	fmt.Println("tocreate: ", permsToCreateIDs)
	err = cluster.AssignMultiplePermissionsToSADB(appCtx, &serviceAccount.ID, permsToCreateIDs)
	if err != nil {
		log.Println(err.Error())
		return nil, status.New(codes.Internal, "Internal error").Err()
	}
	permsToDeleteIDs, err := cluster.UUIDsFromStringArr(permissionsToDelete)
	if err != nil {
		log.Println(err.Error())
		return nil, status.New(codes.InvalidArgument, "invalid permission id").Err()
	}
	fmt.Println("todelete: ", permsToDeleteIDs)
	err = cluster.DeleteMultiplePermissionsOnSADB(appCtx, &serviceAccount.ID, permsToDeleteIDs)
	if err != nil {
		log.Println(err.Error())
		return nil, status.New(codes.Internal, "Internal error").Err()
	}

	// Update single parameters
	serviceAccount.Name = nameRequest
	serviceAccount.Enabled = enabledRequest
	err = cluster.UpdateServiceAccountDB(appCtx, serviceAccount)
	if err != nil {
		log.Println(err.Error())
		return nil, status.New(codes.Internal, "Internal error").Err()
	}

	// if we reach here, all is good, commit
	if err := appCtx.Commit(); err != nil {
		return nil, err
	}

	// get all the updated permissions for the service account
	newPerms, err := cluster.GetAllThePermissionForServiceAccount(appCtx, &serviceAccount.ID)
	if err != nil {
		log.Println(err.Error())
		return nil, status.New(codes.Internal, "Internal error").Err()
	}
	// Build Response
	//transform the permissions to pb format
	var pbPerms []*pb.Permission
	for _, perm := range newPerms {
		pbPerm := buildPermissionPBFromPermission(perm)
		pbPerms = append(pbPerms, pbPerm)
	}
	// update the policies for the SA on Minio
	err = cluster.UpdatePoliciesForSingleServiceAccount(appCtx, &serviceAccount.ID)
	if err != nil {
		log.Println(err.Error())
		return nil, status.New(codes.Internal, "Internal error").Err()
	}

	servAccountsResp := &pb.ServiceAccount{
		Id:        serviceAccount.ID.String(),
		Name:      serviceAccount.Name,
		AccessKey: serviceAccount.AccessKey,
		Enabled:   serviceAccount.Enabled,
	}
	return &pb.InfoServiceAccountResponse{
		ServiceAccount: servAccountsResp,
		Permissions:    pbPerms,
	}, nil
}

func (s *server) InfoServiceAccount(ctx context.Context, in *pb.ServiceAccountActionRequest) (res *pb.InfoServiceAccountResponse, err error) {
	idRequest := in.GetId()
	if idRequest == "" {
		return nil, status.New(codes.InvalidArgument, "an id is needed").Err()
	}
	// start app context
	appCtx, err := cluster.NewTenantContextWithGrpcContext(ctx)
	if err != nil {
		return nil, err
	}
	id, err := uuid.FromString(idRequest)
	serviceAccount, err := cluster.GetServiceAccountByID(appCtx, id)
	if err != nil {
		log.Println(err.Error())
		return nil, status.New(codes.Internal, "Internal error").Err()
	}
	// get all the permissions for the service account
	perms, err := cluster.GetAllThePermissionForServiceAccount(appCtx, &serviceAccount.ID)
	if err != nil {
		log.Println(err.Error())
		return nil, status.New(codes.Internal, "Internal error").Err()
	}

	//transform the permissions to pb format
	var pbPerms []*pb.Permission
	for _, perm := range perms {
		pbPerm := buildPermissionPBFromPermission(perm)
		pbPerms = append(pbPerms, pbPerm)
	}

	servAccountsResp := &pb.ServiceAccount{
		Id:        serviceAccount.ID.String(),
		Name:      serviceAccount.Name,
		AccessKey: serviceAccount.AccessKey,
		Enabled:   serviceAccount.Enabled,
	}
	return &pb.InfoServiceAccountResponse{
		ServiceAccount: servAccountsResp,
		Permissions:    pbPerms,
	}, nil

}
