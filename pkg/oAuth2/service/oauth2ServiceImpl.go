package service

import (
	_oAuth2Repository "github/ariifysp/workshop-go/pkg/oAuth2/repository"
)

type oAuth2ServiceImpl struct {
	oAuth2Repository _oAuth2Repository.OAuth2Repository
}

func NewOAuth2Service(
	oAuth2Repository _oAuth2Repository.OAuth2Repository,
) OAuth2Service {
	return oAuth2ServiceImpl{oAuth2Repository}
}
