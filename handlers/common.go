package handlers

type viewData struct {
	Title    string
	Data     interface{}
	Messages []string
	Errors   []string
}
