## API endpoints

### GET /api/v1/deployment (*DEPRECATED*)

Returns a list of all deployments in the cluster (all namespaces), each with their `name`, `namespace`, `uid`, and `labels`.

#### Response

```json
[
  {
    "name": "deployment-name",
    "namespace": "deployment-namespace",
    "uid": "4ae53bb1-1e42-420f-82a3-a5620ba8b899",
    "labels": {
      "label": "value",
      "label2": "value2"
    }
  }
]
```

### GET /api/v1/pod

Returns a list of all pods in the cluster (all namespaces), each with their `name`, `namespace`, `uid`, `labels` and `ownerReferences`.

Pods belonging to a ReplicaSet will be filtered to only return a single pod per ReplicaSet, using the ReplicaSet's name.

Pods in the `kube-system` namespace will be ignored.

#### Response

```json
[
  {
    "name": "replica-set-name",
    "namespace": "pod-namespace",
    "uid": "4ae53bb1-1e42-420f-82a3-a5620ba8b899",
    "labels": {
      "label": "value",
      "label2": "value2"
    },
    "ownerReferences": [
      {
        "kind": "ReplicaSet",
        "name": "replica-set-name",
        "uid": "9fa96c15-2fbd-4368-8b52-8ef56b437f39"
      }
    ]
  }
]
```

### GET /api/v1/namespace

Returns a list of all namespaces in the cluster, each with their `name`, `uid`, and `labels`.

#### Response

```json
[
  {
    "name": "namespace-name",
    "uid": "4ae53bb1-1e42-420f-82a3-a5620ba8b899",
    "labels": {
      "label": "value",
      "label2": "value2"
    }
  }
]
```

### GET /api/v1/netpol

Returns a list of all NetworkPolicies in the cluster (all namespaces), each with their `metadata` and `spec`.

#### Response

```json
[
  {
    "metadata": {
      "name": "netpol-name",
      "namespace": "netpol-namespace",
    },
    "spec": {
      "podSelector": {
        "matchLabels": {
          "label": "value"
        }
      },
      "policyTypes": [
        "Ingress",
        "Egress"
      ],
      "ingress": [
        {
          "from": [
            {
              "podSelector": {
                "matchLabels": {
                  "label2": "value2"
                }
              }
            }
          ]
        }
      ],
      "egress": [
        {
          "to": [
            {
              "podSelector": {
                "matchLabels": {
                  "label3": "value3"
                }
              }
            }
          ]
        }
      ]
    }
  }
]
```

### POST /api/v1/netpol

Creates a NetworkPolicy in the cluster.

The request body should contain a valid NetworkPolicy including `apiVersion` and `kind` in JSON format.

In case of success, returns the created NetworkPolicy.

#### Request and Response

```json
{
  "apiVersion": "networking.k8s.io/v1",
  "kind": "NetworkPolicy",
  "metadata": {
    "name": "netpol-name",
    "namespace": "netpol-namespace",
  },
  "spec": {
    "podSelector": {
      "matchLabels": {
        "label": "value"
      }
    },
    "policyTypes": [
      "Ingress",
      "Egress"
    ],
    "ingress": [
      {
        "from": [
          {
            "podSelector": {
              "matchLabels": {
                "label2": "value2"
              }
            }
          }
        ]
      }
    ],
    "egress": [
      {
        "to": [
          {
            "podSelector": {
              "matchLabels": {
                "label3": "value3"
              }
            }
          }
        ]
      }
    ]
  }
}
```
