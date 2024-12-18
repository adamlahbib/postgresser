package helpers

import (
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func SetPV(storage string, accessModes []string, id string) *apiv1.PersistentVolume {
	pvAccessModes := make([]apiv1.PersistentVolumeAccessMode, len(accessModes))
	for i, mode := range accessModes {
		pvAccessModes[i] = apiv1.PersistentVolumeAccessMode(mode)
	}
	capacity := apiv1.ResourceList{apiv1.ResourceStorage: resource.MustParse(storage)}
	labelData := map[string]string{
		"app":  "postgres",
		"type": "local",
	}
	return &apiv1.PersistentVolume{
		TypeMeta: metav1.TypeMeta{
			Kind:       "PersistentVolume",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:   PostgresVolumePrefix + id,
			Labels: labelData,
		},
		Spec: apiv1.PersistentVolumeSpec{
			StorageClassName: "manual",
			AccessModes:      pvAccessModes,
			Capacity:         capacity,
			PersistentVolumeSource: apiv1.PersistentVolumeSource{
				HostPath: &apiv1.HostPathVolumeSource{
					Path: "/data/postgresql",
				},
			},
		},
	}
}
