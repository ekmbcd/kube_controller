package models

type Namespace struct {
	Name   string            `json:"name"`
	Uid    string            `json:"uid"`
	Labels map[string]string `json:"labels"`
}
