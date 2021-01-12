package googlesearchconsole

import (
	errortools "github.com/leapforce-libraries/go_errortools"
	google "github.com/leapforce-libraries/go_google"
)

const (
	APIName    string = "GoogleSearchConsole"
	APIURL     string = "https://www.googleapis.com/webmasters/v3"
	DateFormat string = "2006-01-02"
)

// Service stores Service configuration
//
type Service struct {
	googleService *google.Service
}

// methods
//
func NewService(clientID string, clientSecret string, scope string, bigQuery *google.BigQuery) *Service {
	config := google.ServiceConfig{
		APIName:      APIName,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scope:        scope,
	}

	googleService := google.NewService(config, bigQuery)

	return &Service{googleService}
}

func (service *Service) InitToken() *errortools.Error {
	return service.googleService.InitToken()
}
