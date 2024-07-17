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
	Pve_node     string
	Pve_api_user string
	Pve_api_key  string
	Pve_template int
}

type Clone_vm struct {
	Newid   int    `json:"newid"`
	Node    string `json:"node"`
	Name    string `json:"name"`
	Full    bool   `json:"full"`
	Storage string `json:"storage"`
}

func Send_pve_request(http_method string, api_call string, pve_api_creds string, data interface{}) string {

	client := resty.New()
	if http_method == "GET" {
		response, err := client.R().
			EnableTrace().
			SetHeader("Accept", "application/json").
			SetHeader("Authorization", pve_api_creds).
			Get(api_call)
		if err != nil {
			log.Fatal(err)
		}
		return response.Status()

	} else if http_method == "POST" {
		response, err := client.R().
			EnableTrace().
			SetHeader("Accept", "application/json").
			SetHeader("Authorization", pve_api_creds).
			SetBody(data).
			Post(api_call)
		if err != nil {
			log.Fatal(err)
		}
		return response.Status()

	} else if http_method == "PUT" {
		response, err := client.R().
			EnableTrace().
			SetHeader("Accept", "application/json").
			SetHeader("Authorization", pve_api_creds).
			SetBody(data).
			Put(api_call)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("PUT Response:", response.Status())
		return response.Status()

	} else if http_method == "DELETE" {
		response, err := client.R().
			EnableTrace().
			SetHeader("Accept", "application/json").
			SetHeader("Authorization", pve_api_creds).
			SetBody(data).
			Post(api_call)
		if err != nil {
			log.Fatal(err)
		}
		return response.Status()
	} else {
		wrong := "Wrong HTTP Method given!"
		return wrong

	}
	return ""
}

func main() {

	var config = ".config.json"

	content, err := os.ReadFile(config)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var json_payload Json_config
	err = json.Unmarshal(content, &json_payload)

	pve_node := json_payload.Pve_node

	data := Clone_vm{
		Newid:   127,
		Node:    pve_node,
		Name:    "hf-test2",
		Full:    true,
		Storage: "local-storage",
	}

	pve_api_clone := fmt.Sprintf("%s/api2/json/nodes/%s/qemu/%d/clone", json_payload.Pve_host, json_payload.Pve_node, json_payload.Pve_template)
	pve_api_creds := fmt.Sprintf("PVEAPIToken=%s=%s", json_payload.Pve_api_user, json_payload.Pve_api_key)

	clone_vm := Send_pve_request("POST", pve_api_clone, pve_api_creds, data)
	fmt.Println(clone_vm)

}
