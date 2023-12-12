package osprey

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Osprey struct {
	ApiKey string `json:"apiKey"`
	Url    string `json:"Url"`
}

type loggedErr struct {
	ErrorType string `json:"error_type"`
	Message   string `json:"message"`
}

func New(ApiKey string) Osprey {
	formattedUrl := fmt.Sprintf("http://localhost:3000/log?api_key=%s", ApiKey)
	osprey := &Osprey{ApiKey: ApiKey, Url: formattedUrl}
	return *osprey
}

var client http.Client

func (o *Osprey) Log(message string) {

	fmt.Printf("Osprey found a General error. Details: %s\n", message)

	data := loggedErr{ErrorType: "General", Message: message}

	result, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	r, err := http.NewRequest("POST", o.Url, bytes.NewBuffer(result))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-Type", "application/json")

	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
}

func (o *Osprey) Critical(message string) {
	fmt.Printf("Osprey found a Critical error. Details: %s", message)

	data := loggedErr{ErrorType: "Critical", Message: message}

	result, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest("POST", o.Url, bytes.NewBuffer(result))
	if err != nil {
		panic(err)
	}

	request.Header.Add("Content-Type", "application/json")

	res, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

}
