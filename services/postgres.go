package services

import (
	"context"

	helpers "github.com/adamlahbib/postgresser/helpers"
	"github.com/adamlahbib/postgresser/models"
	"github.com/google/uuid"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/retry"
)

type Postgres struct {
	kubeClient *kubernetes.Clientset
}

func NewPostgres(clientset *kubernetes.Clientset) Service {
	return &Postgres{
		kubeClient: clientset,
	}
}

func (p *Postgres) Create(ctx context.Context, request models.CreateRequest) (models.CreateResponse, error) {
	id := uuid.New().String()
	CM := helpers.SetCM(request.DBName, request.Username, request.Password, id)
	PV := helpers.SetPV(request.Capacity, []string{request.AccessMode}, id)
	PVC := helpers.SetPVC(request.Capacity, []string{request.AccessMode}, id)
	deploy := helpers.SetDeploy(request.Replicas, request.Port, id)
	svc := helpers.SetSvc(request.Port, id)
	if _, err := p.kubeClient.CoreV1().ConfigMaps("default").Create(ctx, CM, metav1.CreateOptions{}); err != nil {
		return models.CreateResponse{}, err
	}
	if _, err := p.kubeClient.CoreV1().PersistentVolumes().Create(ctx, PV, metav1.CreateOptions{}); err != nil {
		return models.CreateResponse{}, err
	}
	if _, err := p.kubeClient.CoreV1().PersistentVolumeClaims("default").Create(ctx, PVC, metav1.CreateOptions{}); err != nil {
		return models.CreateResponse{}, err
	}
	if _, err := p.kubeClient.AppsV1().Deployments("default").Create(ctx, deploy, metav1.CreateOptions{}); err != nil {
		return models.CreateResponse{}, err
	}
	if _, err := p.kubeClient.CoreV1().Services("default").Create(ctx, svc, metav1.CreateOptions{}); err != nil {
		return models.CreateResponse{}, err
	}
	return models.CreateResponse{Id: id}, nil
}

func (p *Postgres) Delete(ctx context.Context, request models.DeleteRequest) error {
	deletePolicy := metav1.DeletePropagationForeground // delete pods before deployment
	if err := p.kubeClient.AppsV1().Deployments("default").Delete(ctx, request.Id, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		return err
	}
	if err := p.kubeClient.CoreV1().Services("default").Delete(ctx, request.Id, metav1.DeleteOptions{}); err != nil {
		return err
	}
	if err := p.kubeClient.CoreV1().PersistentVolumeClaims("default").Delete(ctx, request.Id, metav1.DeleteOptions{}); err != nil {
		return err
	}
	if err := p.kubeClient.CoreV1().PersistentVolumes().Delete(ctx, request.Id, metav1.DeleteOptions{}); err != nil {
		return err
	}
	if err := p.kubeClient.CoreV1().ConfigMaps("default").Delete(ctx, request.Id, metav1.DeleteOptions{}); err != nil {
		return err
	}
	return nil
}

func (p *Postgres) Update(ctx context.Context, request models.UpdateRequest) error {
	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// retrieve the latest version of the deployment before attempting update
		// retryonconflict uses exponential backoff to avoid exhausting the apiserver
		result, err := p.kubeClient.AppsV1().Deployments("default").Get(ctx, request.Id, metav1.GetOptions{})
		if err != nil {
			return err
		}
		result.Spec.Replicas = &request.Replicas
		_, err = p.kubeClient.AppsV1().Deployments("default").Update(ctx, result, metav1.UpdateOptions{})
		return err
	})
}
