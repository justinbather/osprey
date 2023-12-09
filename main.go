package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type errLog struct {
	LogType string `json:"logType"`
	Info    string `json:"Info"`
}

func Log(errType string, log string) {
	fmt.Printf("Osprey found a %s, %s\n", errType, log)

	data := errLog{LogType: errType, Info: log}

	result, err := json.Marshal(data)
	if err != nil {

		panic(err)
	}

	r, err := http.NewRequest("POST", "http://localhost:3000/log", bytes.NewBuffer(result))
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

func main() {
	i := 5

	if i != 10 {
		Log("Fatal", "i does not equal 10")
	}

}
