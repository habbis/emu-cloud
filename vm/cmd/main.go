package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"os"
	"strings"
)

type Json_config struct {
	Pve_host     string
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

func Send_request() {

	var config = ".config.json"

	content, err := os.ReadFile(config)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var json_payload Json_config
	err = json.Unmarshal(content, &json_payload)

	// using split to get pve node from url
	pve_host := json_payload.Pve_host
	split_res1 := strings.Split(pve_host, ".")
	res2 := split_res1[0]
	split_res2 := strings.Split(res2, ":")
	res3 := split_res2[1]
	split_res3 := strings.Split(res3, "//")
	pve_node := split_res3[1]

	pve_api_clone := fmt.Sprintf("%s/api2/json/nodes/%s/qemu/%d/clone", json_payload.Pve_host, pve_node, json_payload.Pve_template)
	pve_api_creds := fmt.Sprintf("PVEAPIToken=%s=%s", json_payload.Pve_api_user, json_payload.Pve_api_key)

	data := Clone_vm{
		Newid:   127,
		Node:    pve_node,
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
		Post(pve_api_clone)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("POST Response:", response.Status())

}

func main() {

	var config = ".config.json"

	content, err := os.ReadFile(config)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var json_payload Json_config
	err = json.Unmarshal(content, &json_payload)

	// using split to get pve node from url
	pve_host := json_payload.Pve_host
	split_res1 := strings.Split(pve_host, ".")
	res2 := split_res1[0]
	split_res2 := strings.Split(res2, ":")
	res3 := split_res2[1]
	split_res3 := strings.Split(res3, "//")
	pve_node := split_res3[1]

	pve_api_clone := fmt.Sprintf("%s/api2/json/nodes/%s/qemu/%d/clone", json_payload.Pve_host, pve_node, json_payload.Pve_template)
	pve_api_creds := fmt.Sprintf("PVEAPIToken=%s=%s", json_payload.Pve_api_user, json_payload.Pve_api_key)

	data := Clone_vm{
		Newid:   127,
		Node:    pve_node,
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
		Post(pve_api_clone)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("POST Response:", response.Status())

}
