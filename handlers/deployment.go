package handlers

import (
	"context"
	"log"

	discovery "github.com/gkarthiks/k8s-discovery"
	"github.com/gofiber/fiber/v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"kube-controller/models"
)

func GetDeployments(c *fiber.Ctx, k8s *discovery.K8s) error {

	// get deployments in all the namespaces by omitting namespace
	deployments, err := k8s.Clientset.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error getting deployments %v\n", err)
	}

	filteredDeployments := []models.Deployment{}
	for i := range deployments.Items {
		// remove kube-system deployments
		if deployments.Items[i].Namespace != "kube-system" {
			// transform deployment object into deployment struct
			deploymentStruct := models.Deployment{
				Name:      deployments.Items[i].Name,
				Namespace: deployments.Items[i].Namespace,
				Uid:       string(deployments.Items[i].UID),
				Labels:    deployments.Items[i].Spec.Template.Labels,
			}
			filteredDeployments = append(filteredDeployments, deploymentStruct)
		}
	}

	return c.JSON(filteredDeployments)
}
