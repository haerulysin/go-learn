package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// type IPV4JSON struct

type AWSIPRange struct {
	SyncToken  string `json:"syncToken"`
	CreateDate string `json:"createDate"`
	Prefixes   []struct {
		IPPrefix           string `json:"ip_prefix"`
		Region             string `json:"region"`
		Service            string `json:"service"`
		NetworkBorderGroup string `json:"network_border_group"`
	} `json:"prefixes"`
}

func main() {
	log.SetFlags(0)

	resultPath := filepath.Join(".", "result")
	if _, err := os.Stat(resultPath); os.IsNotExist(err) {
		os.Mkdir(resultPath, os.ModePerm)
	}

	resp, err := http.Get("https://ip-ranges.amazonaws.com/ip-ranges.json")
	if err != nil {
		log.Fatal(err)
	}
	var awsIPRange AWSIPRange
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()
	err = json.Unmarshal(body, &awsIPRange)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile(filepath.Join(resultPath, "result.txt"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for _, objPrefix := range awsIPRange.Prefixes {

		isMyRegion := objPrefix.Region == "ap-southeast-3" || objPrefix.Region == "ap-east-1" || objPrefix.Region == "ap-southeast-1"

		if isMyRegion {
			f.WriteString(objPrefix.IPPrefix + "\n")
		}

	}

}
