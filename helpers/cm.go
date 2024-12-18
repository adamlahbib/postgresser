package helpers

import (
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func SetCM(dbName, user, pass, id string) *apiv1.ConfigMap {
	labelData := map[string]string{
		"app": "postgres",
	}
	postgresData := map[string]string{
		"POSTGRES_DB":       dbName,
		"POSTGRES_USER":     user,
		"POSTGRES_PASSWORD": pass,
	}
	return &apiv1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:   PostgresSecretPrefix + id,
			Labels: labelData,
		},
		Data: postgresData,
	}
}
