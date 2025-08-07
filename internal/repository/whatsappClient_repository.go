package repository

import (
	"context"
	"log"
	"webhook/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type WhatsappClientsRepository struct {
	Collection *mongo.Collection
}

func (wsr *WhatsappClientsRepository) GetClientState(userID string) (*models.WhatsaapClient, error) {
	var chatState models.WhatsaapClient
	filter := bson.M{"user_id": userID}

	err := wsr.Collection.FindOne(context.Background(), filter).Decode(&chatState)
	if err != nil {
		return nil, err
	}

	return &chatState, nil
}

func (wsr *WhatsappClientsRepository) UpdateClientFlow(userID string, newFlow string) error {
	filter := bson.M{"client_id": userID}
	update := bson.M{"$set": bson.M{"current_flow": newFlow}}

	_, err := wsr.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Printf("Error al actualizar el estado del chat: %v", err)
		return err
	}

	log.Printf("Estado del chat para el usuario %s actualizado a: %s\n", userID, newFlow)
	return nil
}

func (wsr *WhatsappClientsRepository) UpdateClientState(userID string, newState string) error {
	filter := bson.M{"client_id": userID}
	update := bson.M{"$set": bson.M{"state": newState}}

	_, err := wsr.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Printf("Error al actualizar el estado del chat: %v", err)
		return err
	}

	log.Printf("Estado del chat para el usuario %s actualizado a: %s\n", userID, newState)
	return nil
}
func (wsr *WhatsappClientsRepository) GetByClientId(clientId string) (*models.WhatsaapClient, error) {
	var client models.WhatsaapClient
	filter := bson.M{"client_id": clientId}
	err := wsr.Collection.FindOne(context.Background(), filter).Decode(&client)
	if err != nil {
		return nil, err
	}

	// El cliente existe
	return &client, err
}

func (wsr *WhatsappClientsRepository) AddClient(client *models.WhatsaapClient) (*mongo.InsertOneResult, error) {
	result, err := wsr.Collection.InsertOne(context.Background(), client)
	if err != nil {
		return nil, err
	}

	return result, nil
}
