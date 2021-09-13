package helpers

type WEBURL struct {
	WebURL    string `json:"web_url"`
	WebStatus string `json:"web_status"`
	Name      string `json:"web_name"`
}

func (w *WEBURL) AddName(name string) {
	w.Name = name
}
