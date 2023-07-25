package handlers

import (
	"context"
	"encoding/json"
	"log"

	discovery "github.com/gkarthiks/k8s-discovery"
	"github.com/gofiber/fiber/v2"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"kube-controller/models"
)

func GetNetpols(c *fiber.Ctx, k8s *discovery.K8s) error {
	// get Network Policies in all the namespaces by omitting namespace
	netpols, err := k8s.Clientset.NetworkingV1().NetworkPolicies("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		log.Printf("Error getting network policies %v\n", err)
	}

	filteredNetpols := []models.NetworkPolicy{}
	for i := range netpols.Items {
		// transform netpol object into netpol struct
		var netpolStruct models.NetworkPolicy
		netpolJson, err := json.Marshal(netpols.Items[i])
		if err != nil {
			log.Printf("Error marshalling netpol %v\n", err)
			return err
		}
		err = json.Unmarshal(netpolJson, &netpolStruct)
		if err != nil {
			log.Printf("Error unmarshalling netpol %v\n", err)
			return err
		}

		filteredNetpols = append(filteredNetpols, netpolStruct)
	}

	return c.JSON(filteredNetpols)
}

func CreateNetpol(c *fiber.Ctx, k8s *discovery.K8s) error {

	netpol := new(networkingv1.NetworkPolicy)

	if err := c.BodyParser(netpol); err != nil {
		log.Printf("Error parsing netpol %v\n", err)
		return err
	}

	netpol, err := k8s.Clientset.NetworkingV1().NetworkPolicies(netpol.Namespace).Create(context.TODO(), netpol, metav1.CreateOptions{})
	if err != nil {
		log.Printf("Error creating netpol %v\n", err)
		return err
	}

	return c.JSON(netpol)
}
