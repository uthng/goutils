package goutils

import (
	"errors"
	//"fmt"
	"encoding/json"
	"strings"
	//"encoding/pem"
	"net/http"

	"crypto/rsa"
	//"crypto/x509"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/spf13/cast"
)

// CustomClaims represents the Azure JWT token
type CustomClaims struct {
	UPN    string   `json:"upn"`
	Groups []string `json:"groups"`

	jwt.StandardClaims
}

// JWTToken represents struct containing signing key or jwks°uri
type JWTToken struct {
	signKey string
	jwksURI string
}

//const (
//azurePublicKeysURL = "https://login.microsoftonline.com/common/discovery/keys"
//)

var (
	// ErrTokenInvalid denotes a token was not able to be validated.
	ErrTokenInvalid = errors.New("JWT Token was invalid")
	// ErrTokenExpired denotes a token'vs expire header (exp) has since passed.
	ErrTokenExpired = errors.New("JWT Token is expired")
	// ErrTokenMalformed denotes a token was not formatted as a JWT token.
	ErrTokenMalformed = errors.New("JWT Token is malformed")
	// ErrTokenNotActive denotes a token's not before header (nbf) is in the
	// future.
	ErrTokenNotActive = errors.New("Token is not valid yet")
	// ErrTokenUnexpectedSigningMethod denotes a token was signed with an unexpected
	// signing method.
	ErrTokenUnexpectedSigningMethod = errors.New("Token unexpected signing method")
	// ErrTokenUnableGetCAPublic denotes not being able to get CA Public Key
	ErrTokenUnableGetCAPublic = errors.New("Token cannot get CA Public Key")
	// ErrTokenKidNotFound denotes token's key identifier not found
	ErrTokenKidNotFound = errors.New("Token's Key Identifier not found")
)

// NewJWTToken returns a new JWTToken
func NewJWTToken() *JWTToken {
	return &JWTToken{}
}

// SetJwksURI sets jwksURI to validate token
func (t *JWTToken) SetJwksURI(uri string) {
	t.jwksURI = uri
}

// SetSignKey set signKey to use when jwksURI is not used.
// Otherwise, it will be computed according to kid & and jwksURI
func (t *JWTToken) SetSignKey(key string) {
	t.signKey = key
}

// ParseToken parses and checks to see if JWT is valid
func (t *JWTToken) ParseToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if token.Method != jwt.SigningMethodRS256 {
			return nil, ErrTokenUnexpectedSigningMethod
		}

		kid := cast.ToString(token.Header["kid"])
		if kid != "" {
			certChain, err := t.getCertChain(kid)
			if err != nil {
				return nil, err
			}

			rsa := cast.ToString(token.Header["alg"])
			if strings.HasPrefix(rsa, "RS") {
				return t.convertPEMToRSAPubKey(certChain)
			}
		}

		return []byte(t.signKey), nil
	})

	if err != nil {
		if e, ok := err.(*jwt.ValidationError); ok {
			switch {
			case e.Errors&jwt.ValidationErrorMalformed != 0:
				return nil, ErrTokenMalformed
			case e.Errors&jwt.ValidationErrorNotValidYet != 0:
				return nil, ErrTokenNotActive
			case e.Errors&jwt.ValidationErrorExpired != 0:
				return nil, ErrTokenExpired
			case e.Errors&jwt.ValidationErrorSignatureInvalid != 0:
				return nil, ErrTokenUnexpectedSigningMethod
			case e.Inner != nil:
				// report e.Inner
				return nil, e.Inner
			}
			// We have a ValidationError but have no specific Go kit error for it.
			// Fall through to return original error.
		}

		return nil, err
	}

	if !token.Valid {
		return nil, ErrTokenInvalid
	}

	//ctx = context.WithValue(ctx, JWTClaimsContextKey, token.Claims)
	return token.Claims.(jwt.MapClaims), nil
}

func (t *JWTToken) getCertChain(kid string) (string, error) {
	var objmap map[string]interface{}

	resp, err := http.Get(t.jwksURI)
	if err != nil {
		return "", err
	}

	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&objmap)
	if err != nil {
		return "", err
	}

	for _, v := range cast.ToSlice(objmap["keys"]) {
		value := cast.ToStringMap(v)
		if cast.ToString(value["kid"]) == kid {
			return cast.ToStringSlice(value["x5c"])[0], nil
		}
	}

	return "", ErrTokenKidNotFound
}

func (t *JWTToken) convertPEMToRSAPubKey(certChain string) (*rsa.PublicKey, error) {
	var PEMSTART = "-----BEGIN CERTIFICATE-----\n"
	var PEMEND = "\n-----END CERTIFICATE-----\n"

	//certPEM := key
	//certPEM = strings.Replace(certPEM, "\\n", "\n", -1)
	//certPEM = strings.Replace(certPEM, "\"", "", -1)
	//block, _ := pem.Decode([]byte(PEMSTART + caPubKey + PEMEND))
	//cert, _ := x509.ParseCertificate(block.Bytes)

	return jwt.ParseRSAPublicKeyFromPEM([]byte(PEMSTART + certChain + PEMEND))
}
