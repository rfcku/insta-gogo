package ig

import (
	"fmt"
	"github.com/rfcku/insta-gogo/utils"
)

type InstagramApi struct {
	AccessToken string
	UserID      string
	ApiBaseURL  string
	Client      *utils.ApiRequestHandler
}

func New(accessToken, userID string) *InstagramApi {
	return &InstagramApi{
		AccessToken: accessToken,
		UserID:      userID,
		ApiBaseURL:  "https://graph.instagram.com",
		Client:      utils.NewApiRequestHandler(),
	}
}

func (api *InstagramApi) WithToken(params map[string]string) map[string]string {
	params["access_token"] = api.AccessToken
	return params
}

func (api *InstagramApi) CreateContainer(imageURL, caption, mediaType string) (string, error) {
	url := fmt.Sprintf("%s/%s/media", api.ApiBaseURL, api.UserID)
	params := api.WithToken(map[string]string{
		"media_type": mediaType,
		"image_url":  imageURL,
		"caption":    caption,
	})
	response, err := api.Client.Post(url, params)

	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	return response["id"].(string), nil

}

func (api *InstagramApi) PublishContainer(containerID, mediaType string) (string, error) {
	url := fmt.Sprintf("%s/%s/media_publish", api.ApiBaseURL, api.UserID)

	params := api.WithToken(map[string]string{
		"media_type":  mediaType,
		"creation_id": containerID,
	})

	response, err := api.Client.Post(url, params)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	return response["id"].(string), nil
}

func (api *InstagramApi) GetUserMedia() (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/media", api.ApiBaseURL, api.UserID)
	params := api.WithToken(map[string]string{})
	return api.Client.Get(url, params)
}

func (api *InstagramApi) CreateAndPublish(imageURL, caption, mediaType string) (string, error) {
	containerID, err := api.CreateContainer(imageURL, caption, mediaType)
	if err != nil {
		return "", err
	}

	publishID, err := api.PublishContainer(containerID, mediaType)
	if err != nil {
		return "", err
	}
	return publishID, nil
}
