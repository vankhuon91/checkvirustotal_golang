package main

import (
	"bufio"
	"checkvirustotal/logfile"
	"checkvirustotal/virstotal"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("---------------------------------------")
	fmt.Println("1. Check hashs: [app] files [hashs.txt]")
	fmt.Println("2. Check domains: [app] domains [domains.txt]")
	fmt.Println("3. Check ips: [app] ip_address [ips.txt]")

	type_search := os.Args[1]
	fileinput := "input.txt"
	switch type_search {
	case "domains":
		fileinput = "domains.txt"
	case "ips":
		fileinput = "ips.txt"
	case "ip_address":
		fileinput = "ip_address.txt"
	}

	logfile.InitLog()

	var lines []string
	file, err := os.Open(fileinput)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed when the function returns

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	// Check for errors during the scan
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	nLines := len(lines)
	fmt.Println("Read file", fileinput)
	fmt.Println("You have ", nLines, "lines", type_search)
	fmt.Print("Start check from line 1 ! Y/n? ")
	editStart := ""
	i := 1
	fmt.Scanln(&editStart)
	if editStart == "n" {
		fmt.Print("Enter start line:")
		fmt.Scanln(&i)
	}
	i = i - 1
	times := 1
	for i < nLines {
		fmt.Println(times, ":", i, "/", nLines, lines[i])
		err := virstotal.Search(type_search, lines[i], times)
		if err != nil {
			//fmt.Println(err.Error())
			time.Sleep(2 * time.Second)
		} else {
			i++
		}
		times++
	}
}
