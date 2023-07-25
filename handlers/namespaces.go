package handlers

import (
	"context"
	"encoding/json"
	"kube-controller/models"
	"log"

	discovery "github.com/gkarthiks/k8s-discovery"
	"github.com/gofiber/fiber/v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetNamespaces(c *fiber.Ctx, k8s *discovery.K8s) error {

	namespaces, err := k8s.Clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error getting namespaces %v\n", err)
		return err
	}

	filteredNamespaces := []models.Namespace{}
	for i := range namespaces.Items {
		// remove kube-system namespaces
		if namespaces.Items[i].Name != "kube-system" {
			// transform namespace object into namespace struct
			var namespaceStruct models.Namespace
			namespaceJson, err := json.Marshal(namespaces.Items[i].ObjectMeta)
			if err != nil {
				log.Printf("Error marshalling namespace %v\n", err)
				return err
			}
			err = json.Unmarshal(namespaceJson, &namespaceStruct)
			if err != nil {
				log.Printf("Error unmarshalling namespace %v\n", err)
				return err
			}
			filteredNamespaces = append(filteredNamespaces, namespaceStruct)
		}
	}

	return c.JSON(filteredNamespaces)
}
