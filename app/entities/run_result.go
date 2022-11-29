package entities

type RunResult struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
}
