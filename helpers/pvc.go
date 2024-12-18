package helpers

import (
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func SetPVC(storage string, accessModes []string, id string) *apiv1.PersistentVolumeClaim {
	pvcAccessModes := make([]apiv1.PersistentVolumeAccessMode, len(accessModes))
	for i, mode := range accessModes {
		pvcAccessModes[i] = apiv1.PersistentVolumeAccessMode(mode)
	}
	labelData := map[string]string{
		"app": "postgres",
	}
	storageClassName := "manual"
	capacity := apiv1.ResourceList{apiv1.ResourceStorage: resource.MustParse(storage)}
	return &apiv1.PersistentVolumeClaim{
		TypeMeta: metav1.TypeMeta{
			Kind:       "PersistentVolumeClaim",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:   PostgresVolumeClaimPrefix + id,
			Labels: labelData,
		},
		Spec: apiv1.PersistentVolumeClaimSpec{
			AccessModes:      pvcAccessModes,
			StorageClassName: &storageClassName,
			Resources: apiv1.VolumeResourceRequirements{
				Requests: capacity,
			},
		},
	}
}
