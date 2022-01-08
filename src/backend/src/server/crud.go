package server

import (
	"encoding/json"
	"errors"
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

	// then we basically do a last check, if it's well-formed response, add in cors + other stuff
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	defer func() {
		if err != nil {
			log.Printf("error occured whilee processing request! %v", err)
		}
	}()

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
	} else if method == http.MethodPut {
		err = s.handlePutRequest(w, r, body)
		return
	}

	http.Error(w, "method not supported", http.StatusMethodNotAllowed)
	err = errors.New("unsupported method")
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

func (s *Server) handleGetRequest(w http.ResponseWriter, r *http.Request, body []byte) error {
	if len(body) == 0 {
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

	// then we get the item based on id
	getBody, err := s.getIdFromBody(body)
	if err != nil {
		http.Error(w, "malformed body", http.StatusBadRequest)
		return err
	}

	item, err := s.GetItemFromId(r.Context(), getBody.Id)
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

func (s *Server) handleDeleteRequest(w http.ResponseWriter, r *http.Request, body []byte) error {
	if len(body) == 0 {
		// let's not let them delete everything, so we'll just do this as nothing, error
		http.Error(w, "id not specified", http.StatusBadRequest)
		return errors.New("id not specified")
	}

	// then we get the item based on id
	getBody, err := s.getIdFromBody(body)
	if err != nil {
		http.Error(w, "malformed body", http.StatusBadRequest)
		return err
	}

	err = s.DeleteItemFromId(r.Context(), getBody.Id)
	if err != nil {
		http.Error(w, "could not delete item...", http.StatusInternalServerError)	
		return err
	}

	w.WriteHeader(http.StatusOK)
	return nil
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

func (s *Server) handlePutRequest(w http.ResponseWriter, r *http.Request, body []byte) error {
	type data struct {
		Id string
		Data string
	}

	var getBody data
	err := json.Unmarshal(body, &getBody)

	if err != nil {
		log.Printf("Error parsing body into struct: %v", err)
		http.Error(w, "malformed body", http.StatusBadRequest)
		return err
	}

	err = s.UpdateItem(r.Context(), getBody.Id, getBody.Data)
	if err != nil {
		http.Error(w, "could not create new item...", http.StatusInternalServerError)	
		return err
	}

	w.WriteHeader(http.StatusOK)
	return nil
}
