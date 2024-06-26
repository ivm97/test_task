package models

type Action struct {
	Title  string            `json:"name"`
	Result string            `json:"result"`
	Params map[string]string `json:"params"`
}

type Data struct {
	Actions []Action `json:"actions"`
}
