package common

type Secret struct {
	Encrypted string `json:"encrypted,omitempty"`
	Plain     string `json:"plain,omitempty"`
}

type PortRange struct {
	Deny  bool `json:"deny,omitempty"`
	End   int  `json:"end,omitempty"`
	Start int  `json:"start,omitempty"`
}
