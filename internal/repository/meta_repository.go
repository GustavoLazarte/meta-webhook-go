package repository

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Repository para manejar la descarga del media
type MetaRepository struct {
}

func (repo *MetaRepository) DownloadMedia(mediaID string) ([]byte, error) {
	url := fmt.Sprintf("https://graph.facebook.com/v19.0/%s", mediaID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("META_TOKEN"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error al descargar media: %s", resp.Status)
	}

	mediaData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return mediaData, nil
}
