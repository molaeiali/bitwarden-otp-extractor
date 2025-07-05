package main

import (
	"bitwarden-otp-extractor/aegis"
	"bitwarden-otp-extractor/bitwarden"
	"bitwarden-otp-extractor/utils"
	"encoding/json"
	"fmt"
	"os"
)

func printUsage(supportedModes  []string) {
	usage := "Usage: bitwarden-otp-extractor <mode> <bitwarden export file address> [output file address]"
	fmt.Println(usage)
	fmt.Println("Supported modes:",supportedModes)
}

func main(){
	supportedModes := []string{"aegis"}

	if(len(os.Args) == 1){
		printUsage(supportedModes)
		return
	}
	
	if (!utils.ContainsString(supportedModes, os.Args[1])) {
		printUsage(supportedModes)
		fmt.Println("Mode not supported")
		return
	}
	mode := os.Args[1]
	
	if len(os.Args) < 3 {
		printUsage(supportedModes)
		fmt.Println("Bitwarden export file address is needed")
		return
	}

	outPath := mode + ".json"

	if len(os.Args) == 4 {
		outPath = os.Args[3]
	}

	file, err := os.Open(os.Args[2])

	if(err != nil) {
		fmt.Println(err)
		return
	}

	defer file.Close()

	var b bitwarden.Bitwarden
	
	decoder := json.NewDecoder(file)

	err = decoder.Decode(&b) 
	if(err != nil) {
		fmt.Println(err)
	}

	var json []byte

	if(mode == "aegis"){
		json, err = aegis.Convert(b);
		if(err != nil) {
			fmt.Println(err)
			return
		}
	}

	err = os.WriteFile(outPath, json, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

}
