package models

type Pod struct {
	Name            string            `json:"name"`
	Namespace       string            `json:"namespace"`
	Uid             string            `json:"uid"`
	Labels          map[string]string `json:"labels"`
	OwnerReferences []OwnerReference  `json:"ownerReferences"`
}

type OwnerReference struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	UID  string `json:"uid"`
}
