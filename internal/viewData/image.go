package viewData

import (
	"fmt"
	"encoding/json"
	"net/http"
)

const pixabayAPIKey = "43064707-905dc35f16b4f01b61a0a08e6"
const pixabayBaseURL = "https://pixabay.com/api/"

type PixabayResponse struct {
    Hits []struct {
        WebformatURL string `json:"webformatURL"`
    } `json:"hits"`
}

func SearchImages(query string) ([]string, error) {
    url := fmt.Sprintf("%s?key=%s&q=%s", pixabayBaseURL, pixabayAPIKey, query)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var pixabayResp PixabayResponse
    if err := json.NewDecoder(resp.Body).Decode(&pixabayResp); err != nil {
        return nil, err
    }

    var imageURLs []string
    for _, hit := range pixabayResp.Hits {
        imageURLs = append(imageURLs, hit.WebformatURL)
		break
    }
    return imageURLs, nil
}
