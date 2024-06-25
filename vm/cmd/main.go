package main

import (
	_ "encoding/json"
	_ "fmt"
	"github.com/go-resty/resty/v2"
	"log"
	_ "log"
	_ "os"
)

func Get_vm_list() {

	// Create a Resty Client
	client := resty.New()

	// Sample of using Request.SetQueryString method
	resp, err := client.R().
		//SetQueryString("productId=232&template=fresh-sample&cat=resty&source=google&kw=buy a lot more").
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", "PVEAPIToken=temp=temp").
		Get("https://pve:8006/api2/json'")

	if err != nil {
		log.Fatal(err)
	}

	println(resp)

}

func main() {
	err := Get_vm_list

	if err != nil {
		log.Fatal(err)
	}

}
