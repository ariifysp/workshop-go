package controller

import (
	_oAuth2Service "github/ariifysp/workshop-go/pkg/oAuth2/service"
)

type oAuth2ControllerImpl struct {
	oAuth2Service _oAuth2Service.OAuth2Service
}

func NewOAuth2Controller(
	oAuth2Service _oAuth2Service.OAuth2Service,
) OAuth2Controller {
	return &oAuth2ControllerImpl{oAuth2Service}
}
