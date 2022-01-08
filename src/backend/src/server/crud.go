package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (s *Server) crud(w http.ResponseWriter, r *http.Request, id string) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	method := r.Method
	// not the best design, but since it's crud, we probably don't need to add more than this...
	if method == http.MethodGet {
		s.handleSingleGetRequest(w, r, body, id)
		return
	} else if method == http.MethodDelete {
		s.handleSingleDeleteRequest(w, r, body, id)
		return
	} else if method == http.MethodPut {
		s.handlePutRequest(w, r, body, id)
		return
	}

	http.Error(w, "method not supported", http.StatusMethodNotAllowed)
}

func (s *Server) handleSingleGetRequest(w http.ResponseWriter, r *http.Request, body []byte, id string) error {
	item, err := s.GetItemFromId(r.Context(), id)
	if err != nil {
		http.Error(w, "can't get items from db", http.StatusInternalServerError)
		return err
	}

	returnBody, err := json.Marshal(item)
	if err != nil {
		log.Printf("Error returning body: %v", err)
		http.Error(w, "can't return body", http.StatusInternalServerError)
		return err
	}

	w.Write(returnBody)
	return nil
}

func (s *Server) handleSingleDeleteRequest(w http.ResponseWriter, r *http.Request, body []byte, id string) error {
	err := s.DeleteItemFromId(r.Context(), id)
	if err != nil {
		http.Error(w, "could not delete item...", http.StatusInternalServerError)	
		return err
	}

	w.WriteHeader(http.StatusOK)
	return nil
}

func (s *Server) handlePutRequest(w http.ResponseWriter, r *http.Request, body []byte, id string) error {
	type data struct {
		Data string
	}

	var getBody data
	err := json.Unmarshal(body, &getBody)

	if err != nil {
		log.Printf("Error parsing body into struct: %v", err)
		http.Error(w, "malformed body", http.StatusBadRequest)
		return err
	}

	err = s.UpdateItem(r.Context(), id, getBody.Data)
	if err != nil {
		http.Error(w, "could not create new item...", http.StatusInternalServerError)	
		return err
	}

	w.WriteHeader(http.StatusOK)
	return nil
}
