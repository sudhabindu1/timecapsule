package model

type Message struct {
	Sender       string `json:"sender"`
	Body         string `json:"body"`
	RevealMillis int64  `json:"reveal_millis,omitempty"`
}

func (m Message) IsValid() bool {
	return len(m.Body) <= 280 && m.Sender != "" // && time.Now().UnixNano()/1000000 < m.RevealMillis
}
