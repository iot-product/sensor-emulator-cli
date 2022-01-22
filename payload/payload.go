package payload

type Payload struct {
	Id        int    `json:id`
	Data      int    `json:data`
	Timestamp string `json:timestamp`
}
