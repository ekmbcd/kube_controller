package models

type Pod struct {
	Name            string            `json:"name"`
	Namespace       string            `json:"namespace"`
	Uid             string            `json:"uid"`
	Labels          map[string]string `json:"labels,omitempty"`
	OwnerReferences []OwnerReference  `json:"ownerReferences,omitempty"`
}

type OwnerReference struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	UID  string `json:"uid"`
}
