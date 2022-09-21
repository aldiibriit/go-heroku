package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	port := "7177"
	handler := func(w http.ResponseWriter, req *http.Request) {
		i := 1
		response := ""
		for {
			if i > 99 {
				break
			}
			multiplication := `99 x ` + strconv.Itoa(i) + ` = ` + strconv.Itoa(99*i) + ``
			response += multiplication
			response += "\n"
			i++
		}
		io.WriteString(w, response)
	}

	http.HandleFunc("/api/multiplication99", handler)
	log.Println("Listing for :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
