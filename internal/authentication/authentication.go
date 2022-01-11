package authentication

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

type Service struct {
	// database instance
}

type Authentication struct {
	Username string
	Password string
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

type AuthenticationService interface {
	CreateJWTToken(a *Authentication) (string, error)
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CreateJWTToken(a *Authentication) (*TokenDetails, error) {

	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.NewV4().String()

	var err error
	//Creating Access Token
	td.AccessToken, err = createAccessToken(a.Username)
	if err != nil {
		return nil, err
	}

	//Creating Refresh Token
	td.RefreshToken, err = createRefreshToken(a.Username, td)
	if err != nil {
		return nil, err
	}

	return td, nil
}

func createAccessToken(username string) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = username
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func createRefreshToken(username string, td *TokenDetails) (string, error) {
	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf") //this should be in an env file
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] =  username
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	refreshToken, err := rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return "", err
	}
	return refreshToken, err
}
