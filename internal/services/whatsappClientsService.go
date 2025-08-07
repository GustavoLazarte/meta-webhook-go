package services

import (
	"webhook/internal/models"
	"webhook/internal/models/requests"
	"webhook/internal/models/responses"
	"webhook/internal/repository"
	"webhook/internal/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

type WhatsappClients struct {
	*repository.WhatsappClientsRepository
}

func (wcs *WhatsappClients) FindByClientId(id string) (bool, *responses.WhatsappClient, error) {
	var clientResponse responses.WhatsappClient
	client, err := wcs.WhatsappClientsRepository.GetByClientId(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil, nil
		}

		return false, nil, err
	}

	err = utils.MapStruct(client, clientResponse)
	if err != nil {
		return false, nil, err
	}
	return true, &clientResponse, nil
}

func (wcs *WhatsappClients) CreateNewClient(id string) error {
	var newClientModel models.WhatsaapClient
	newClient := &requests.WhatsappClient{
		ClientID:    id,
		CurrentFlow: "",
		State:       "new",
	}

	err := utils.MapStruct(newClient, &newClientModel)
	if err != nil {
		return err
	}

	_, err = wcs.WhatsappClientsRepository.AddClient(&newClientModel)
	if err != nil {
		return err
	}

	return nil
}

func (wcs *WhatsappClients) UpdateClientFlow(id string, currentFlow string) error {
	err := wcs.WhatsappClientsRepository.UpdateClientFlow(id, currentFlow)
	if err != nil {
		return err
	}
	return nil
}
