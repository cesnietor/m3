// This file is part of MinIO Kubernetes Cloud
// Copyright (c) 2020 MinIO, Inc.
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

package restapi

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/minio/m3/cluster"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/minio/m3/models"
	"github.com/minio/m3/restapi/operations"
	"github.com/minio/m3/restapi/operations/admin_api"
)

func registerStoragePoolsHandlers(api *operations.M3API) {
	// get StoragePools info
	api.AdminAPIGetStoragePoolsInfoHandler = admin_api.GetStoragePoolsInfoHandlerFunc(func(params admin_api.GetStoragePoolsInfoParams, principal *models.Principal) middleware.Responder {
		sessionID := string(*principal)
		resp, err := getStoragePoolsInfoResponse(sessionID)
		if err != nil {
			return admin_api.NewGetStoragePoolsInfoDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return admin_api.NewGetStoragePoolsInfoOK().WithPayload(resp)

	})
}

func getStoragePoolsInfo(ctx context.Context, client K8sClient) (*models.StoragePoolsInfo, error) {
	dat, err := os.Open("fakeapi/storagepools.json")

	if err != nil {
		log.Println("Error Opening Test-StoragePools File:", err)
	}
	defer dat.Close()
	byteValue, _ := ioutil.ReadAll(dat)
	var spInfo *models.StoragePoolsInfo

	json.Unmarshal(byteValue, &spInfo)

	return spInfo, nil
}

func getStoragePoolsInfoResponse(token string) (*models.StoragePoolsInfo, error) {
	ctx := context.Background()
	client, err := cluster.K8sClient(token)
	if err != nil {
		log.Println("error getting k8sClient:", err)
		return nil, err
	}
	k8sClient := &k8sClient{
		client: client,
	}
	storagePoolsInfo, err := getStoragePoolsInfo(ctx, k8sClient)
	if err != nil {
		log.Println("error getting storage pools info:", err)
		return nil, err

	}
	return storagePoolsInfo, nil
}
