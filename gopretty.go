package gopretty

import (
	"encoding/json"
	"fmt"
	"log"
)

// Print anything in a JSON fashion
func Print(v ...interface{}) {
	for _, value := range v {
		b, err := json.MarshalIndent(value, "", "  ")
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(string(b))
	}
}
