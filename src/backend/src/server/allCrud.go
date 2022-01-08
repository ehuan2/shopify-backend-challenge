package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (s *Server) allCrud(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	defer func() {
		if err != nil {
			log.Printf("error occured while processing request! %v", err)
		}
	}()

	method := r.Method

	// not the best design, but since it's crud, we probably don't need to add more than this...
	if method == http.MethodGet {
		err = s.handleGetRequest(w, r, body)
		return
	} else if method == http.MethodDelete {
		err = s.handleDeleteRequest(w, r, body)
		return
	} else if method == http.MethodPost {
		err = s.handlePostRequest(w, r, body)
		return
	}
}

func (s *Server) handleGetRequest(w http.ResponseWriter, r *http.Request, body []byte) error {
	// then we get all the items
	items, err := s.GetAllItems(r.Context())
	if err != nil {
		http.Error(w, "can't get items from db", http.StatusInternalServerError)
		return err
	}

	returnBody, err := json.Marshal(items)
	if err != nil {
		log.Printf("Error returning body: %v", err)
		http.Error(w, "can't return body", http.StatusInternalServerError)
		return err
	}
	w.Write(returnBody)
	return nil
}

func (s *Server) handleDeleteRequest(w http.ResponseWriter, r *http.Request, body []byte) error {
	return s.DeleteAllItems(r.Context())
}

func (s *Server) handlePostRequest(w http.ResponseWriter, r *http.Request, body []byte) error {
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

	newId, err := s.CreateNewItem(r.Context(), getBody.Data)
	if err != nil {
		http.Error(w, "could not create new item...", http.StatusInternalServerError)	
		return err
	}

	type newItemResponse struct {
		Id string
	}
	newItem:= newItemResponse {
		Id: newId,
	}

	returnBody, err := json.Marshal(newItem)
	if err != nil {
		log.Printf("Error returning body: %v", err)
		http.Error(w, "can't return body", http.StatusInternalServerError)
		return err
	}

	w.Write(returnBody)
	return nil
}

