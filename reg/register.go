package reg

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	contentType = "application/json"
)

type Registration struct {
	Service   string     `json:"service"`
	Version   string     `json:"version"`
	Instances []Instance `json:"instances"`
}

type Instance struct {
	Location string      `json:"location"`
	Capacity interface{} `json:"capacity"`
}

func Send(path string, location string) {
	log.Println("Reading in registration file.")
	r := readRegJSON(path)
	log.Println("Sending registration.")
	if regJson, err := json.Marshal(r); err != nil {
		log.Println("Error creating json for request", err)
		os.Exit(1)
	} else {
		if res, err := http.Post(location, contentType, bytes.NewReader(regJson)); err != nil {
			log.Println("Error sending registration: ", err)
			os.Exit(1)
		} else {
			log.Println("Response from elsa: ", res)
		}
	}
}

func readRegJSON(path string) Registration {
	var reg Registration
	if file, err := ioutil.ReadFile(path); err != nil {
		log.Println("Error reading registration file: ", err)
		os.Exit(1)
	} else {
		if err = json.Unmarshal(file, &reg); err != nil {
			log.Println("Error parsing registration json: ", err)
			os.Exit(1)
		}
	}
	return reg
}
