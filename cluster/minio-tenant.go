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

package cluster

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// mkTenantContainer builds the container spec for the `sgTenant` tenant on host identified by `hostNum`
func mkTenantMinioContainer(sgTenant *StorageGroupTenant, sgNode *StorageGroupNode) (v1.Container, []v1.Volume) {
	sg := sgTenant.StorageGroup
	envName := fmt.Sprintf("%s-env", sgTenant.ShortName)
	volumeMounts := []v1.VolumeMount{}
	tenantContainer := v1.Container{
		Name:  fmt.Sprintf("%s-minio-%d", sgTenant.Tenant.ShortName, sgNode.Num),
		Image: "minio/minio:RELEASE.2019-10-12T01-39-57Z",
		//Image:           "minio/minio:latest",
		ImagePullPolicy: "IfNotPresent",
		Args: []string{
			"server",
			"--address",
			fmt.Sprintf(":%d", sgTenant.Port),
			fmt.Sprintf(
				"http://sg-%d-host-{1...%d}:%d/mnt/tdisk{1...%d}",
				sg.Num,
				sgTenant.StorageGroup.TotalNodes,
				sgTenant.Port,
				sgTenant.StorageGroup.TotalVolumes),
		},
		Ports: []v1.ContainerPort{
			{
				Name:          "http",
				ContainerPort: sgTenant.Port,
			},
		},
		EnvFrom: []v1.EnvFromSource{
			{
				SecretRef: &v1.SecretEnvSource{
					LocalObjectReference: v1.LocalObjectReference{Name: envName},
				},
			},
		},
		LivenessProbe: &v1.Probe{
			Handler: v1.Handler{
				HTTPGet: &v1.HTTPGetAction{
					Path: "/minio/health/live",
					Port: intstr.IntOrString{
						IntVal: sgTenant.Port,
					},
				},
			},
			InitialDelaySeconds: 120,
			PeriodSeconds:       20,
		},
	}
	var tenantVolumes []v1.Volume
	//volumes that will be used by this sgTenant
	for _, vol := range sgNode.Node.Volumes {
		vname := fmt.Sprintf("%s-pv-%d", sgTenant.Tenant.ShortName, vol.Num)
		volumenSource := v1.VolumeSource{HostPath: &v1.HostPathVolumeSource{Path: fmt.Sprintf("%s/%s", vol.MountPath, sgTenant.Tenant.ShortName)}}
		hostPathVolume := v1.Volume{Name: vname, VolumeSource: volumenSource}

		mount := v1.VolumeMount{
			Name:      vname,
			MountPath: fmt.Sprintf("/mnt/tdisk%d", vol.Num),
		}
		tenantVolumes = append(tenantVolumes, hostPathVolume)
		volumeMounts = append(volumeMounts, mount)
	}
	tenantContainer.VolumeMounts = volumeMounts

	return tenantContainer, tenantVolumes
}
