package helpers

import (
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func SetDeploy(replicas, port int32, id string) *appsv1.Deployment {
	matchLabels := map[string]string{"app": "postgres"}
	labels := map[string]string{"app": "postgres"}
	return &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: id,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: matchLabels,
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: apiv1.PodSpec{
					Volumes: []apiv1.Volume{{
						Name: "postgresdata",
						VolumeSource: apiv1.VolumeSource{
							PersistentVolumeClaim: &apiv1.PersistentVolumeClaimVolumeSource{
								ClaimName: PostgresVolumeClaimPrefix + id,
							},
						},
					}},
					Containers: []apiv1.Container{{
						Name:            "postgres",
						Image:           "postgres:latest",
						ImagePullPolicy: "IfNotPresent",
						Ports: []apiv1.ContainerPort{{
							ContainerPort: port,
						}},
						EnvFrom: []apiv1.EnvFromSource{{
							ConfigMapRef: &apiv1.ConfigMapEnvSource{
								LocalObjectReference: apiv1.LocalObjectReference{
									Name: PostgresSecretPrefix + id,
								},
							},
						}},
						VolumeMounts: []apiv1.VolumeMount{{
							Name:      "postgresdata",
							MountPath: "/var/lib/postgresql/data",
						}},
					}},
				},
			},
		},
	}
}
