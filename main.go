/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Note: the example only works with the code within the same release/branch.
package main

import (
	"flag"
	"kube-controller/handlers"
	"log"

	discovery "github.com/gkarthiks/k8s-discovery"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

func main() {
	k8s, err := discovery.NewK8s()
	if err != nil {
		log.Fatal(err)
	}

	port := flag.String("port", ":3000", "Port to listen on")
	flag.Parse()

	app := fiber.New()

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Create a /api/v1 endpoint
	v1 := app.Group("/api/v1")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Miao 🐱")
	})

	v1.Get("/deployment", func(c *fiber.Ctx) error {
		return handlers.GetDeployments(c, k8s)
	})

	v1.Get("/netpol", func(c *fiber.Ctx) error {
		return handlers.GetNetpols(c, k8s)
	})

	// Listen on port 3000
	log.Fatal(app.Listen(*port)) // go run app.go -port=:3000

}
