package handlers

import (
	"context"
	"encoding/json"
	"log"

	discovery "github.com/gkarthiks/k8s-discovery"
	"github.com/gofiber/fiber/v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"kube-controller/models"
)

func GetPods(c *fiber.Ctx, k8s *discovery.K8s) error {
	pods, err := k8s.Clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error getting pods %v\n", err)
	}

	filteredPods := []models.Pod{}
	for i := range pods.Items {
		// remove kube-system pods
		if pods.Items[i].Namespace != "kube-system" {
			// consider replicasets as single pods
			if (pods.Items[i].OwnerReferences) != nil {
				alreadyIn := false
				for _, pod := range filteredPods {
					if pod.OwnerReferences[0].Name == pods.Items[i].OwnerReferences[0].Name {
						alreadyIn = true
					}
				}
				if alreadyIn {
					continue
				}
			}

			// transform pod object into pod struct
			var podStruct models.Pod
			podJson, err := json.Marshal(pods.Items[i].ObjectMeta)
			if err != nil {
				log.Printf("Error marshalling pod %v\n", err)
				return err
			}
			err = json.Unmarshal(podJson, &podStruct)
			if err != nil {
				log.Printf("Error unmarshalling pod %v\n", err)
				return err
			}

			filteredPods = append(filteredPods, podStruct)
		}
	}

	return c.JSON(filteredPods)
}
