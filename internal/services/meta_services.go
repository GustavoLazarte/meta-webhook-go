package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"webhook/internal/factory/messages_factory"
	"webhook/internal/messages"
	"webhook/internal/repository"
	"webhook/internal/utils"
)

type MetaService struct {
	repository repository.MetaRepository
}

func (service *MetaService) DownloadAndSaveMedia(mediaID, fileName string) error {
	// Descarga el media usando el repository
	mediaData, err := service.repository.DownloadMedia(mediaID)
	if err != nil {
		return fmt.Errorf("Error al descargar el media: %v", err)
	}

	// Guardar el media en un archivo
	if err := saveToFile(fileName, mediaData); err != nil {
		return fmt.Errorf("Error al guardar el archivo: %v", err)
	}

	log.Printf("Archivo descargado y guardado exitosamente en: %s", fileName)
	return nil
}

func saveToFile(fileName string, data []byte) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("Error al crear el archivo: %v", err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("Error al escribir los datos en el archivo: %v", err)
	}

	return nil
}

func (service *MetaService) SendInteractiveMessage(to string) error {
	message := messages_factory.BuildMessageInteractive(to)
	messageBody, _ := json.Marshal(message)
	body := bytes.NewBuffer(messageBody)
	req, err := utils.CreateRequest(body)
	if err != nil {
		return err
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error procesando el mensaje al evniar: %v", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error procesando el mensaje al recibir: %v", err)
		return err
	}

	return nil
}

func (service *MetaService) SendTextMessage(to string, text string) error {
	message := messages_factory.BuildMessageText(to, text)
	messageBody, _ := json.Marshal(message)
	body := bytes.NewBuffer(messageBody)
	req, err := utils.CreateRequest(body)
	if err != nil {
		return err
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error procesando el mensaje al evniar: %v", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error procesando el mensaje al recibir: %v", err)
		return err
	}

	return nil
}

func (service *MetaService) SendTemplateMessage(templateName string, to string, param []string) error {
	message := messages_factory.BuildMessageTemplate(templateName, to, param[0], param[1])
	messageBody, _ := json.Marshal(message)
	body := bytes.NewBuffer(messageBody)
	req, err := utils.CreateRequest(body)
	if err != nil {
		return err
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error procesando el mensaje al evniar: %v", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Lee el cuerpo de la respuesta para obtener más detalles sobre cualquier error.
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error leyendo el cuerpo de la respuesta: %v", err)
			return err
		}

		// Convierte el cuerpo a string para loguearlo.
		bodyString := string(bodyBytes)
		log.Printf("Error procesando el mensaje. Status code: %d, Body: %s", resp.StatusCode, bodyString)

		return fmt.Errorf("error en la respuesta: status code %d, body: %s", resp.StatusCode, bodyString)
	}

	return nil
}

func (service *MetaService) SendTemplateMessageV2(template *messages.TemplateBody) error {
	message := template
	messageBody, _ := json.Marshal(message)
	body := bytes.NewBuffer(messageBody)
	req, err := utils.CreateRequest(body)
	if err != nil {
		return err
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error procesando el mensaje al evniar: %v", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Lee el cuerpo de la respuesta para obtener más detalles sobre cualquier error.
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error leyendo el cuerpo de la respuesta: %v", err)
			return err
		}

		// Convierte el cuerpo a string para loguearlo.
		bodyString := string(bodyBytes)
		log.Printf("Error procesando el mensaje. Status code: %d, Body: %s", resp.StatusCode, bodyString)

		return fmt.Errorf("error en la respuesta: status code %d, body: %s", resp.StatusCode, bodyString)
	}

	return nil
}

func (service *MetaService) SetRepository(metaRepository repository.MetaRepository) {
	service.repository = metaRepository
}
