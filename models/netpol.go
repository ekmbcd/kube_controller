package models

type PolicyType string

const (
	Ingress PolicyType = "Ingress"
	Egress  PolicyType = "Egress"
)

type Selector struct {
	MatchLabels      map[string]string `json:"matchLabels,omitempty"`
	MatchExpressions []MatchExpression `json:"matchExpressions,omitempty"`
}

type Port struct {
	Protocol string `json:"protocol"`
	Port     int    `json:"port"`
	EndPort  int    `json:"endPort,omitempty"`
}

type IPBlock struct {
	CIDR   string   `json:"cidr,omitempty"`
	Except []string `json:"except,omitempty"`
}

type MatchExpression struct {
	Key      string   `json:"key"`
	Operator string   `json:"operator"`
	Values   []string `json:"values"`
}

type Policy struct {
	PodSelector       *Selector `json:"podSelector,omitempty"`
	IPBlock           *IPBlock  `json:"ipBlock,omitempty"`
	NamespaceSelector *Selector `json:"namespaceSelector,omitempty"`
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
		PodSelector Selector        `json:"podSelector"`
		PolicyTypes []PolicyType    `json:"policyTypes"`
		Ingress     []IngressPolicy `json:"ingress,omitempty"`
		Egress      []EgressPolicy  `json:"egress,omitempty"`
	} `json:"spec"`
}
