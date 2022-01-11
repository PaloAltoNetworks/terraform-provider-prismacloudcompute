package common

type Secret struct {
	Encrypted string `json:"encrypted"`
	Plain     string `json:"plain"`
}
