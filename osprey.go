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

// Prototype function to work with temporary server
func (o *Osprey) Log(errorType string, message string) {
	fmt.Printf("Osprey found a %s error, %s\n", errorType, message)

	data := loggedErr{ErrorType: errorType, Message: message}

	result, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(result))

	//Temp api key so we can track errors for an account

	r, err := http.NewRequest("POST", o.Url, bytes.NewBuffer(result))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
}
