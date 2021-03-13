package googlesearchconsole

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	google "github.com/leapforce-libraries/go_google"
	bigquery "github.com/leapforce-libraries/go_google/bigquery"
)

const (
	apiName    string = "GoogleSearchConsole"
	apiURL     string = "https://www.googleapis.com/webmasters/v3"
	dateFormat string = "2006-01-02"
)

// Service stores Service configuration
//
type Service struct {
	googleService *google.Service
}

// methods
//
func NewService(clientID string, clientSecret string, scope string, bigQueryService *bigquery.Service) *Service {
	config := google.ServiceConfig{
		APIName:      apiName,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scope:        scope,
	}

	googleService := google.NewService(config, bigQueryService)

	return &Service{googleService}
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("%s/%s", apiURL, path)
}

func (service *Service) InitToken() *errortools.Error {
	return service.googleService.InitToken()
}
