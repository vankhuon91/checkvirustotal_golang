package virstotal

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/VirusTotal/vt-go"
)

type Resultvt struct {
	Malicious  int `json:"malicious"`
	Suspicious int `json:"suspicious"`
	Undetected int `json:"undetected"`
	Harmless   int `json:"harmless"`
	Timeout    int `json:"timeout"`
}

var apikeys = []string{"f0b3881151830154668814748625c2284b451ce41ef93b6dba92d6492dd145f4"}

func Search(type_str, str string, key int) error {
	nkey := key % len(apikeys)
	client := vt.NewClient(apikeys[nkey])

	data, err := client.GetObject(vt.URL("%s/%s", type_str, str))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	analys_interface, _ := data.Get("last_analysis_stats")
	analys_byte, _ := json.Marshal(analys_interface)
	analys_struct := Resultvt{}
	err_parse := json.Unmarshal(analys_byte, &analys_struct)
	if err_parse != nil {
		fmt.Println(err_parse.Error())
		return err_parse
	}
	log.Println(str, ".Malicious", analys_struct.Malicious, ".Totals", analys_struct.Harmless+analys_struct.Malicious+analys_struct.Undetected+analys_struct.Suspicious)

	return nil
}
