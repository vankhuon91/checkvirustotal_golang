package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	vt "github.com/VirusTotal/vt-go"
)

type resultvt struct {
	Malicious  string `json:"malicious"`
	Suspicious string `json:"suspicious"`
	Undetected string `json:"undetected"`
	Harmless   string `json:"harmless"`
	Timeout    string `json:"timeout"`
}

func main() {
	fmt.Println("--------------Start check virustotal-------------")
	apikeys := []string{"f0b3881151830154668814748625c2284b451ce41ef93b6dba92d6492dd145f4"}
	client := vt.NewClient(apikeys[0])
	args := os.Args[1:]
	fmt.Println(args)
	data, err := client.GetObject(vt.URL("%s/%s", args[0], args[1]))
	if err != nil {
		log.Println(err.Error())
	}
	analys_interface, _ := data.Get("last_analysis_stats")
	analys_byte, _ := json.Marshal(analys_interface)
	analys_struct, _ := json.Unmarshal(analys_byte)
	fmt.Println(analys_struct)
}
