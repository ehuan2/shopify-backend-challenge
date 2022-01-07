package server

import (
	"context"
	"log"

	"github.com/google/uuid"
)

// this file here manages the database operations, in order to be able to choose which db to use in the future
// essentially decoupling the server's handling of http requests and the actual database fetching
// again, not the best design, but since we'll probably not change this much, makes sense to decouple
type Item struct {
	Id uuid.UUID
	Data string
}

// change these implementations if we want to change the db we use
func (s *Server) GetAllItems(ctx context.Context) ([]Item, error) {
	keys, err := s.redisDb.Keys(ctx, "*").Result()
	if err != nil {
		log.Printf("Could not get all keys %v", err)
		return nil, err
	}

	var items []Item

	for _, key := range keys {
		item, err := s.GetItemFromId(ctx, key)
		if err != nil || item == nil { // nil check the item
			log.Printf("Could not get item with id: %v", err)
			return nil, err
		}
		items = append(items, *item)
	}

	return items, nil
}

func (s *Server) GetItemFromId(ctx context.Context, key string) (*Item, error) {
	value, err := s.redisDb.Get(ctx, key).Result()
	if err != nil {
		log.Printf("Could not get value from key: %v", err)
		return nil, err
	}
	id, err := uuid.Parse(key)
	if err != nil {
		log.Printf("Could not parse key: %v", err)
		return nil, err
	}
	return &Item{
		Id: id,
		Data: value,
	}, nil
}

func (s *Server) DeleteItemFromId(ctx context.Context, key string) error {
	err := s.redisDb.Del(ctx, key).Err()
	if err != nil {
		log.Printf("Could not delete item with key: %v", err)
		return err
	}
	return nil
}

func (s *Server) CreateNewItem(ctx context.Context, data string) (string, error) {
	newUUId := uuid.New()
	err := s.redisDb.Set(ctx, newUUId.String(), data, 0).Err()
	if err != nil {
		log.Printf("Could not create new item: %v", err)
		return "", err
	}
	return newUUId.String(), nil
}

func (s *Server) UpdateItem(ctx context.Context, key string, data string) error {
	err := s.redisDb.Set(ctx, key, data, 0).Err()
	if err != nil {
		log.Printf("Could not update item: %v", err)
		return err
	}
	return nil
}

