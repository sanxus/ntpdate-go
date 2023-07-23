package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func main() {
	server := os.Args[1]

	// send our request
	resp, err := ntp.Query(server)
	if err != nil {
		panic(err)
	}
	// validate the response
	err = resp.Validate()
	if err != nil {
		panic(err)
	}

	if contains(os.Args, "-json") {
		json, err := json.Marshal(resp)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(json))
	} else {
		fmt.Printf("server %s, stratum %d, offset %+.6f, delay %.5f\n",
			server, resp.Stratum, resp.ClockOffset.Seconds(), resp.RTT.Seconds())
		fmt.Printf("Now %v\n", resp.Time.Local())
	}
}
