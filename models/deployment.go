package models

type Deployment struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Uid       string            `json:"uid"`
	Ip        string            `json:"ip"`
	Labels    map[string]string `json:"labels"`
}
