package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (s *Server) crud(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	method := r.Method
	// not the best design, but since it's crud, we probably don't need to add more than this...
	if method == http.MethodGet {
		s.handleGetRequest(w, r, body)
		return
	} else if method == http.MethodDelete {
		s.handleDeleteRequest(w, r, body)
		return
	} else if method == http.MethodPost {
		s.handlePostRequest(w, r, body)
		return
	} else if method == http.MethodPut {
		s.handlePutRequest(w, r, body)
		return
	}

	http.Error(w, "method not supported", http.StatusMethodNotAllowed)
}

// otherwise, we attempt to read the body and get an id from it
type requestWithId struct {
	Id string
}

func (s *Server) getIdFromBody(body []byte) (*requestWithId, error) {
	// then we unmarshal it to the proper id
	var getBody *requestWithId
	err := json.Unmarshal(body, &getBody)
	if err != nil {
		log.Printf("Error parsing body into struct: %v", err)
		return nil, err
	}
	return getBody, nil
}

func (s *Server) handleGetRequest(w http.ResponseWriter, r *http.Request, body []byte) {
	if len(body) == 0 {
		// then we get all the items
		items, err := s.GetAllItems(r.Context())
		if err != nil {
			http.Error(w, "can't get items from db", http.StatusInternalServerError)
			return
		}

		returnBody, err := json.Marshal(items)
		if err != nil {
			log.Printf("Error returning body: %v", err)
			http.Error(w, "can't return body", http.StatusInternalServerError)
			return
		}
		w.Write(returnBody)
		return
	}

	// then we get the item based on id
	getBody, err := s.getIdFromBody(body)
	if err != nil {
		http.Error(w, "malformed body", http.StatusBadRequest)
		return
	}

	item, err := s.GetItemFromId(r.Context(), getBody.Id)
	if err != nil {
		http.Error(w, "can't get items from db", http.StatusInternalServerError)
		return
	}

	returnBody, err := json.Marshal(item)
	if err != nil {
		log.Printf("Error returning body: %v", err)
		http.Error(w, "can't return body", http.StatusInternalServerError)
		return
	}
	w.Write(returnBody)
}

func (s *Server) handleDeleteRequest(w http.ResponseWriter, r *http.Request, body []byte) {
	if len(body) == 0 {
		// let's not let them delete everything, so we'll just do this as nothing, error
		http.Error(w, "id not specified", http.StatusBadRequest)
		return
	}

	// then we get the item based on id
	getBody, err := s.getIdFromBody(body)
	if err != nil {
		http.Error(w, "malformed body", http.StatusBadRequest)
		return
	}

	err = s.DeleteItemFromId(r.Context(), getBody.Id)
	if err != nil {
		http.Error(w, "could not delete item...", http.StatusInternalServerError)	
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) handlePostRequest(w http.ResponseWriter, r *http.Request, body []byte) {
	type data struct {
		Data string 
	}
	var getBody data
	err := json.Unmarshal(body, &getBody)

	if err != nil {
		log.Printf("Error parsing body into struct: %v", err)
		http.Error(w, "malformed body", http.StatusBadRequest)
		return
	}

	newId, err := s.CreateNewItem(r.Context(), getBody.Data)
	if err != nil {
		http.Error(w, "could not create new item...", http.StatusInternalServerError)	
		return
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
		return
	}

	w.Write(returnBody)
}

func (s *Server) handlePutRequest(w http.ResponseWriter, r *http.Request, body []byte) {
	type data struct {
		Id string
		Data string
	}

	var getBody data
	err := json.Unmarshal(body, &getBody)

	if err != nil {
		log.Printf("Error parsing body into struct: %v", err)
		http.Error(w, "malformed body", http.StatusBadRequest)
		return
	}

	err = s.UpdateItem(r.Context(), getBody.Id, getBody.Data)
	if err != nil {
		http.Error(w, "could not create new item...", http.StatusInternalServerError)	
		return
	}

	w.WriteHeader(http.StatusOK)
}
