package aegis

import (
	"bitwarden-otp-extractor/bitwarden"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type Aegis struct {
	Version int64 `json:"version"`
	Header AegisHeader `json:"header"`
	Db AegisDb `json:"db"`
}

type AegisHeader struct {
	Slots *string `json:"slots"`
	Params *string `json:"params"`
}

type AegisDb struct {
	Version int64 `json:"version"`
	Entries []AegisEntry `json:"entries"`
}

type AegisEntryInfo struct {
	Secret string `json:"secret"`
	Algo	string `json:"algo"`
	Digits string `json:"digits"`
	Period string `json:"period"`
}

type AegisEntry struct {
	Type string `json:"type"`
	Uuid string `json:"uuid"`
	Name string `json:"name"`
	Issuer string `json:"issuer"`
	Note string `json:"note"`
	Favorite bool `json:"favorite"`
	Icon *string `json:"icon"`
	Info AegisEntryInfo `json:"info"`
	Groups []string `json:"groups"`
}

func Convert(b bitwarden.Bitwarden) ([]byte, error){
	var a Aegis
	a.Version = 1
	a.Header.Params = nil
	a.Header.Slots = nil
	a.Db.Version = 3
	
	for _, item := range b.Items {
		if(item.Login.Totp != "") {
			uriString := item.Login.Totp

			parsedURL, err := url.Parse(uriString)
			if err != nil {
				fmt.Printf("Error parsing URI: %v\n", err)
				return nil, err
			}

			label := parsedURL.Path   

			if len(label) > 0 && label[0] == '/' {
				label = label[1:]
			}

			queryParams := parsedURL.Query()

			secret := queryParams.Get("secret")
			algorithm := queryParams.Get("algorithm")
			digitsStr := queryParams.Get("digits")
			periodStr := queryParams.Get("period")
			issuer := queryParams.Get("issuer")

			if(issuer == ""){
				issuer = label
			}

			digits := 0
			if digitsStr != "" {
				digits, err = strconv.Atoi(digitsStr)
				if err != nil {
					fmt.Printf("Error converting digits to int: %v\n", err)
					digits = 6 
				}
			} else {
				digits = 6 
			}


			period := 0
			if periodStr != "" {
				period, err = strconv.Atoi(periodStr)
				if err != nil {
					fmt.Printf("Error converting period to int: %v\n", err)
					period = 30
				}
			} else {
				period = 30
			}

			var entry AegisEntry
			entry.Favorite = false
			entry.Issuer = issuer
			entry.Name = item.Name
			entry.Note = ""
			entry.Icon = nil
			entry.Type = "totp"
			entry.Uuid = item.Id
			entry.Info.Secret = secret
			entry.Info.Algo = algorithm
			entry.Info.Digits = strconv.Itoa(digits)
			entry.Info.Period = strconv.Itoa(period)
			entry.Groups = []string{}
			a.Db.Entries = append(a.Db.Entries, entry)
		}
	}
	
	return json.MarshalIndent(a, "", "	")
}