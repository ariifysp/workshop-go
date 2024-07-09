package repository

type oAuth2RepositoryImpl struct{}

func NewOAuth2RepositoryImpl() OAuth2Repository {
	return &oAuth2RepositoryImpl{}
}
