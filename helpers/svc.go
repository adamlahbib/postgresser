package helpers

import (
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func SetSvc(port int32, id string) *apiv1.Service {
	labelData := map[string]string{
		"app": "postgres",
	}
	selector := map[string]string{
		"app": "postgres",
	}
	return &apiv1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:   PostgresPrefix + id,
			Labels: labelData,
		},
		Spec: apiv1.ServiceSpec{
			Ports:    []apiv1.ServicePort{{Port: port}},
			Selector: selector,
			Type:     "NodePort",
		},
	}
}
