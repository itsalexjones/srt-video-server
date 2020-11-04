package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/haivision/srtgo"
)

func main() {
	options := make(map[string]string)
	options["blocking"] = "0"
	options["transtype"] = "live"
	options["latency"] = "350"
	hostname := "cloudlayer.itsalexjones.com"
	port := 5060

	log.Printf("(srt://%s:%d) Calling", hostname, port)
	a := srtgo.NewSrtSocket(hostname, uint16(port), options)
	err := a.Connect()
	defer a.Close()
	if err != nil {
		panic("Error on Connect")
	}

	go fetchData(a)

	for {
		time.Sleep(500 * time.Millisecond)
		stats, err := a.Stats()
		if err != nil {
			log.Printf("Error in stats: %s", err.Error())
		}

		data, err := json.Marshal(stats)
		if err != nil {
			log.Printf("Erorr json converting stats: %s", err.Error())
		}

		fmt.Printf("%s\n", data)
	}

}

func fetchData(a *srtgo.SrtSocket) {
	buff := make([]byte, 2048)
	for {
		n, err := a.Read(buff, 10000)

		if err != nil {
			fmt.Println(err)
			break
		}

		if n == 0 {
			break
		}

	}
}
