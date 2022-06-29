package customrules

type Ref struct {
	Id     int    `json:"_id,omitempty"`
	Action string `json:"action,omitempty"`
	Effect string `json:"effect,omitempty"`
}
