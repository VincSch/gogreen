package main

import (
	"log"
	"net/http"
	"time"
	"fmt"
	"encoding/json"
	"github.com/rs/cors"
)

var measumrents []Measurement

func main() {
	measumrents = append(measumrents,
		Measurement{Time: time.Now(), Temperature: &Temperature{Time: time.Now(), Value: 36.9},
			Ph:           &Ph{Time: time.Now(), Value: 7.9}})
	measumrents = append(measumrents,
		Measurement{Time: time.Now(), Temperature: &Temperature{Time: time.Now(), Value: 31.9},
			Ph:           &Ph{Time: time.Now(), Value: 3.9}})
	measumrents = append(measumrents,
		Measurement{Time: time.Now(), Temperature: &Temperature{Time: time.Now(), Value: 23.9},
			Ph:           &Ph{Time: time.Now(), Value: 4.9}})
	measumrents = append(measumrents,
		Measurement{Time: time.Now(), Temperature: &Temperature{Time: time.Now(), Value: 28.9},
			Ph:           &Ph{Time: time.Now(), Value: 5.6}})

	DeleteMeasurements()

	for _, item := range measumrents {
		StoreMeasurement(item)
		asJson, _ := json.Marshal(item)
		fmt.Println(string(asJson))
	}

	handler := cors.Default().Handler(RestAPI())
	log.Fatal(http.ListenAndServe(":12345", handler))
}
