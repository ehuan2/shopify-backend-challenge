package csv

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type csvData struct {
	Id string
	Data struct {
		Metadata string
		Type string
		Cost float64
	}
}

func (s *Server) csv(w http.ResponseWriter, r *http.Request) {
	// here we'll get all the inventory by calling the server
	resp, err := http.Get("http://backend:8080/items/") // should be found through backend container
	if err != nil {
		log.Printf("Error getting items: %v", err)
		http.Error(w, "could not get items...", http.StatusInternalServerError)
		return
	}

	// read the body into bytes
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	// marshal it into the csv data array	
	var allData []csvData

	err = json.Unmarshal(body, &allData)
	if err != nil {
		log.Printf("Could not unmarshal data: %v", err)
		http.Error(w, "internal error while getting data", http.StatusInternalServerError)
		return
	}

	// first set necessary headers, for csv exports + cors
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Disposition", "attachment; filename=export.csv")
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Transfer-Encoding", "chunked")

	// then we'll loop through all the entries and write to the response
	writer := csv.NewWriter(w)
	err = writer.Write([]string{"Id", "Metadata", "Type", "Cost"})
	if err != nil {
		log.Printf("could not write to file: %v", err)
		http.Error(w, "internal error while writing to file", http.StatusInternalServerError)
		return
	}

	for _, entry := range allData {
		err = writer.Write([]string{entry.Id, entry.Data.Metadata, entry.Data.Type, fmt.Sprintf("%v", entry.Data.Cost)})
		if err != nil {
			log.Printf("could not write to file: %v", err)
			http.Error(w, "internal error while writing to file", http.StatusInternalServerError)
			return
		}
	}

	writer.Flush()

}
