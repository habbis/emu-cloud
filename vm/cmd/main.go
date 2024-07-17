package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"os"
)

type Json_config struct {
	Pve_host     string
	Pve_api_user string
	Pve_api_key  string
	Pve_tempalte int
}

type Clone_vm struct {
	Newid   int    `json:"newid"`
	Node    string `json:"node"`
	Name    string `json:"name"`
	Full    bool   `json:"full"`
	Storage string `json:"storage"`
}

func main() {

	var config = ".config.json"

	content, err := os.ReadFile(config)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var json_payload Json_config
	err = json.Unmarshal(content, &json_payload)

	pve_clone_api := fmt.Sprintf("%s/api2/json/nodes/hf-pve7/qemu/%d/clone", json_payload.Pve_tempalte, json_payload.Pve_tempalte)
	pve_api_creds := fmt.Sprintf("PVEAPIToken=%s=%s", json_payload.Pve_api_user, json_payload.Pve_api_key)

	data := Clone_vm{
		Newid:   127,
		Node:    "hf-pve7",
		Name:    "hf-test2",
		Full:    true,
		Storage: "local-storage",
	}

	client := resty.New()

	response, err := client.R().
		EnableTrace().
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", pve_api_creds).
		SetBody(&data).
		Post(pve_clone_api)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("POST Response:", response.Status())

}
