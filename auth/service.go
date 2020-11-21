package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

//https://github.com/dgrijalva/jwt-go

type Service interface {
	GenerateToken(userID int) (string, error) //Param user ID yang akan diparsing ke JWT dan kembaliannya berupa string dari Token
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte("bwabwabwabwabwabwabwaakusayangkamu4ever7h3ul71m473")

func NewServiceAuth() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{} //Deklarasi Payloadnya
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) //Generate Token dengan algoritma HS256

	//Setiap JWT dianggap valid jika Secret key request sama dengan secret key pada server(ditandatangani)
	signedToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(encodedToken *jwt.Token) (interface{}, error) {
		_, ok := encodedToken.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid Token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
