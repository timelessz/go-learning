package main

import "fmt"
import "encoding/json"
import "bytes"

type Track struct {
	XmlRequest string `json:"xmlRequest"`
}

func (t *Track) JSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

func main() {
	message := Track{}
	message.XmlRequest = "<car><mirror>XML</mirror></car>"
	fmt.Println("Before Marshal", message)
	messageJSON, _ := message.JSON()
	fmt.Println("After marshal", string(messageJSON))
}
