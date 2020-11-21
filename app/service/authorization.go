package service

import (
	"api-listeners/app/util"
	"time"
)

type AuthorizationService interface {
	GetAuthorizationToken() (string, error)
}

type JwtAuthorizationService struct {
	LoginUrl string
	Username string
	Password string
	TokenTimeToLiveMinutes int64
	authzTokenCache authorizationTokenCache
}

type authorizationTokenCache struct {
	authzToken string
	issuedAt time.Time
}

func (service *JwtAuthorizationService) GetAuthorizationToken() (string, error) {
	if service.hasValidCachedToken() {
		return service.authzTokenCache.authzToken, nil
	}
	authzToken, err := service.requestNewJwt()
	if err != nil {
		return "", err
	}
	service.authzTokenCache = authorizationTokenCache{authzToken: authzToken, issuedAt: time.Now()}
	return authzToken, nil
}

func (service *JwtAuthorizationService) hasValidCachedToken() bool {
	if service.authzTokenCache.authzToken != "" {
		timeTokenLives := time.Now().Sub(service.authzTokenCache.issuedAt)
		maxTokenTTL := time.Duration(service.TokenTimeToLiveMinutes) * time.Minute
		return timeTokenLives < maxTokenTTL
	}
	return false
}

func (service *JwtAuthorizationService) requestNewJwt() (string, error) {
	authReq := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{Username: service.Username, Password: service.Password}
	authResp := struct {
		Response struct {
			Data struct {
				Token string `json:"token"`
			} `json:"data"`
		} `json:"response"`
	}{}
	err := util.DoPostJson(service.LoginUrl, authReq, &authResp)
	if err != nil {
		return "", err
	}
	return authResp.Response.Data.Token, nil
}