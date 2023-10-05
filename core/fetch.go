package core

import (
	"net/http"

	"github.com/spf13/viper"
)

func SetTifAuthHeaders(req *http.Request) {
	req.Header.Set("Ocp-Apim-Subscription-Key", viper.GetString("ApiKey"))
	req.Header.Set("token", viper.GetString("AccessToken"))
	req.Header.Set("x-api-key", viper.GetString("x-api-key"))
}
