package virstotal

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/VirusTotal/vt-go"
)

type Resultvt struct {
	Malicious  int `json:"malicious"`
	Suspicious int `json:"suspicious"`
	Undetected int `json:"undetected"`
	Harmless   int `json:"harmless"`
	Timeout    int `json:"timeout"`
}

var apikeys = []string{
	"ed85db0a7394ae12758ebb936f5b5dacf1b8d585187221f48ceee2cbea969d11",
	"c36c86c315d12e4e3e1eb633800e2b71ea57d6eb5c82ea8470c320b352bf78c5",
	"d143408f7014997dce82b21b76cece8bf3f262155c3c0b7e1ae865280cf21bd6",
	"07fe7526e70d2fd1b7dc352ab1c51eeb8b6f18889d2d367378150ea81a267aae",
	"42f38aed6e9ade2489b09e07645df8f917897a856cd96b5e81fb7e1479976b38",
	"45bb837de0e8fc2bc1f26b423118d4398eacb3c74abd44a47e0b3707749e3453",
	"68682200d46f23511d5e6c0e6371125403f06b7e3b4fbcf0055bdc53d3c2334c",
	"ce6641d0679795022c07bb9519196c7918a01b5f97fe1f725667f12ab1f359df"}

func Search(type_str, str string, line_no int, times int) error {
	nkey := times % len(apikeys)
	client := vt.NewClient(apikeys[nkey])

	data, err := client.GetObject(vt.URL("%s/%s", type_str, str))
	if err != nil {
		fmt.Println(err.Error())
		if (err.Error() == "Resource not found.") || (strings.Contains(err.Error(), "valid")) {
			log.Println("Error:", err.Error(), "Data", str)
			return nil
		}
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
	log.Println("STT:", line_no, "Data:", str, "Malicious", analys_struct.Malicious, "Totals", analys_struct.Harmless+analys_struct.Malicious+analys_struct.Undetected+analys_struct.Suspicious)

	return nil
}
