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

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/spf13/cast"
)

// JWTToken represents struct containing signing key or jwks°uri
type JWTToken struct {
	signMethod string
	signKey    string
	jwksURI    string
}

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
// It will be used to sign/verify the token.
func (t *JWTToken) SetSignKey(key string) {
	t.signKey = key
}

// SetSignMethod set signing method to sign or verify the token
func (t *JWTToken) SetSignMethod(method string) {
	t.signMethod = method
}

// ParseToken parses and checks to see if JWT is valid.
// If the token is valid, it returns a map of all claims defined
// in the JWT.
func (t *JWTToken) ParseToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if token.Method != jwt.GetSigningMethod(t.signMethod) {
			return nil, ErrTokenUnexpectedSigningMethod
		}

		// If key identifier is defined in the header,
		// we try to find out the algo and jwks_uri (x5c)
		// and get the corresponding signing public key.
		// If the algo is RSA, we also need to convert PEM format
		// to RSA pub key.
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

		// If key identifier is not defined, we use signKey instead.
		return []byte(t.signKey), nil
	})

	// Check token'ts validity according to some criterons.
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
