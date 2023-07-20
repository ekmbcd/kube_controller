package models

type PolicyType string

const (
	Ingress PolicyType = "Ingress"
	Egress  PolicyType = "Egress"
)

type PodSelector struct {
	MatchLabels map[string]string `json:"matchLabels,omitempty"`
}

type Port struct {
	Protocol string `json:"protocol"`
	Port     int    `json:"port"`
}

type IPBlock struct {
	CIDR   string   `json:"cidr,omitempty"`
	Except []string `json:"except,omitempty"`
}

type NamespaceSelector struct {
	MatchLabels map[string]string `json:"matchLabels,omitempty"`
}

type Policy struct {
	PodSelector       *PodSelector       `json:"podSelector,omitempty"`
	IPBlock           *IPBlock           `json:"ipBlock,omitempty"`
	NamespaceSelector *NamespaceSelector `json:"namespaceSelector,omitempty"`
}

type IngressPolicy struct {
	Ports []Port   `json:"ports"`
	From  []Policy `json:"from,omitempty"`
}

type EgressPolicy struct {
	Ports []Port   `json:"ports"`
	To    []Policy `json:"to,omitempty"`
}

type NetworkPolicy struct {
	Metadata struct {
		Name      string `json:"name"`
		Namespace string `json:"namespace"`
	} `json:"metadata"`
	Spec struct {
		PodSelector PodSelector     `json:"podSelector"`
		PolicyTypes []PolicyType    `json:"policyTypes"`
		Ingress     []IngressPolicy `json:"ingress"`
		Egress      []EgressPolicy  `json:"egress"`
	} `json:"spec"`
}
