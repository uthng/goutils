package goutils_test

import (
	//"encoding/json"
	//"fmt"
	//"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/uthng/goutils"
)

func TestParseToken(t *testing.T) {

	testCases := []struct {
		name       string
		token      string
		signMethod string
		signKey    string
		jwksURI    string
		result     interface{}
	}{
		{
			"ErrTokenExpired",
			"eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6IllNRUxIVDBndmIwbXhvU0RvWWZvbWpxZmpZVSIsImtpZCI6IllNRUxIVDBndmIwbXhvU0RvWWZvbWpxZmpZVSJ9.eyJhdWQiOiJodHRwczovL21hbmFnZW1lbnQuY29yZS53aW5kb3dzLm5ldC8iLCJpc3MiOiJodHRwczovL3N0cy53aW5kb3dzLm5ldC82MTI5OTBhNC1mZWJkLTQxYWItOWNkMC02ZmU4ODRjYzhiZWUvIiwiaWF0IjoxNTg0MDk3Njk2LCJuYmYiOjE1ODQwOTc2OTYsImV4cCI6MTU4NDEwMTU5NiwiYWNyIjoiMSIsImFpbyI6IkFWUUFxLzhPQUFBQXdjMDVtSk1HWmtOYU5IU0RsSnhxTGVrNFlsc3hxdXBhRGdMNGtlbzYvbGZxcVRDQWFSTmY3S1VaTll1MDhZUkFBNFMwbzM5Q0J4TXFyRHBOZFg3LzdOeGVJakZPQmxHTDRmcHhBZUZuM0FRPSIsImFtciI6WyJwd2QiLCJtZmEiXSwiYXBwaWQiOiIwNGIwNzc5NS04ZGRiLTQ2MWEtYmJlZS0wMmY5ZTFiZjdiNDYiLCJhcHBpZGFjciI6IjAiLCJncm91cHMiOlsiMmU0MDlhMGQtMzZlNy00YjU3LWJmMWEtOWM0MjgzOWQ0NjQyIiwiZTliZThhMGItZDc0OS00YTVjLWJlYmEtNzRmZmNmMDU5YTc0IiwiMDU1ZjczYmItYmJlYS00NmE1LWI0Y2QtYzExODAyYmQyNDgwIl0sImlwYWRkciI6IjE5NS42OC41MC4yMjYiLCJuYW1lIjoiVGhhbmggTmd1eWVuIiwib2lkIjoiZjQ1NWY0NWEtMzRlOC00ZDk1LTllNDUtODk3MTFkYmU1NWM4IiwicHVpZCI6IjEwMDMyMDAwNzkxMjVFRkMiLCJzY3AiOiJ1c2VyX2ltcGVyc29uYXRpb24iLCJzdWIiOiJIYU1qSDNLRmVad3RPMmx1cHZKU05GZnVMZDlRZzNFbzkyOGRnNVlLNnNNIiwidGlkIjoiNjEyOTkwYTQtZmViZC00MWFiLTljZDAtNmZlODg0Y2M4YmVlIiwidW5pcXVlX25hbWUiOiJ0aGFuaC5uZ3V5ZW5AZGl0dG9jbG91ZG91dGxvb2tjb20ub25taWNyb3NvZnQuY29tIiwidXBuIjoidGhhbmgubmd1eWVuQGRpdHRvY2xvdWRvdXRsb29rY29tLm9ubWljcm9zb2Z0LmNvbSIsInV0aSI6InQ4TDRxOFhEVlU2bmJvX1c2X3FwQUEiLCJ2ZXIiOiIxLjAiLCJ3aWRzIjpbImNmMWMzOGU1LTM2MjEtNDAwNC1hN2NiLTg3OTYyNGRjZWQ3YyIsIjYyZTkwMzk0LTY5ZjUtNDIzNy05MTkwLTAxMjE3NzE0NWUxMCJdfQ.kX3uDtQRORUYcpzV3ZmuMxn62_iYlxomOhzJzhITZdcvTmleoBPoXaQ6uxqWpzQhJn2y_1GyVcD9cObGhRW1wlQ8zb0QIQiEUcmMfZ7LmFRjG33-EsVu6lW2boc1Oec_XIze57Yzuiy1mR_l1aqbHAbji4Q-PLhY4uBtNq6ZHQY2vTafg2t2VbmxczmHw10YXKHwqYC7bmS68i_yEO2x0iOr1XK-uRqsXWOLJGgv4UIrm1SmJX7IVC0WRu-OC_AfRUWHzELdY27CEaE9tMHkRkPJOETY3nMvDRxE8VKpQ1AW2aRFGr2GvL_S1m3AFXeCPl1XI9KyrmXWOfpPTIzCLw",
			"RS256",
			"",
			"https://login.microsoftonline.com/common/discovery/keys",
			goutils.ErrTokenKidNotFound,
		},
		{
			"ErrTokenMalformed",
			"WpxZmpZVSIsImtpZCI6IllNRUxIVDBndmIwbXhvU0RvWWZvbWpxZmpZVSJ9.eyJhdWQiOiJodHRwczovL21hbmFnZW1lbnQuY29yZS53aW5kb3dzLm5ldC8iLCJpc3MiOiJodHRwczovL3N0cy53aW5kb3dzLm5ldC82MTI5OTBhNC1mZWJkLTQxYWItOWNkMC02ZmU4ODRjYzhiZWUvIiwiaWF0IjoxNTg0MDk3Njk2LCJuYmYiOjE1ODQwOTc2OTYsImV4cCI6MTU4NDEwMTU5NiwiYWNyIjoiMSIsImFpbyI6IkFWUUFxLzhPQUFBQXdjMDVtSk1HWmtOYU5IU0RsSnhxTGVrNFlsc3hxdXBhRGdMNGtlbzYvbGZxcVRDQWFSTmY3S1VaTll1MDhZUkFBNFMwbzM5Q0J4TXFyRHBOZFg3LzdOeGVJakZPQmxHTDRmcHhBZUZuM0FRPSIsImFtciI6WyJwd2QiLCJtZmEiXSwiYXBwaWQiOiIwNGIwNzc5NS04ZGRiLTQ2MWEtYmJlZS0wMmY5ZTFiZjdiNDYiLCJhcHBpZGFjciI6IjAiLCJncm91cHMiOlsiMmU0MDlhMGQtMzZlNy00YjU3LWJmMWEtOWM0MjgzOWQ0NjQyIiwiZTliZThhMGItZDc0OS00YTVjLWJlYmEtNzRmZmNmMDU5YTc0IiwiMDU1ZjczYmItYmJlYS00NmE1LWI0Y2QtYzExODAyYmQyNDgwIl0sImlwYWRkciI6IjE5NS42OC41MC4yMjYiLCJuYW1lIjoiVGhhbmggTmd1eWVuIiwib2lkIjoiZjQ1NWY0NWEtMzRlOC00ZDk1LTllNDUtODk3MTFkYmU1NWM4IiwicHVpZCI6IjEwMDMyMDAwNzkxMjVFRkMiLCJzY3AiOiJ1c2VyX2ltcGVyc29uYXRpb24iLCJzdWIiOiJIYU1qSDNLRmVad3RPMmx1cHZKU05GZnVMZDlRZzNFbzkyOGRnNVlLNnNNIiwidGlkIjoiNjEyOTkwYTQtZmViZC00MWFiLTljZDAtNmZlODg0Y2M4YmVlIiwidW5pcXVlX25hbWUiOiJ0aGFuaC5uZ3V5ZW5AZGl0dG9jbG91ZG91dGxvb2tjb20ub25taWNyb3NvZnQuY29tIiwidXBuIjoidGhhbmgubmd1eWVuQGRpdHRvY2xvdWRvdXRsb29rY29tLm9ubWljcm9zb2Z0LmNvbSIsInV0aSI6InQ4TDRxOFhEVlU2bmJvX1c2X3FwQUEiLCJ2ZXIiOiIxLjAiLCJ3aWRzIjpbImNmMWMzOGU1LTM2MjEtNDAwNC1hN2NiLTg3OTYyNGRjZWQ3YyIsIjYyZTkwMzk0LTY5ZjUtNDIzNy05MTkwLTAxMjE3NzE0NWUxMCJdfQ.kX3uDtQRORUYcpzV3ZmuMxn62_iYlxomOhzJzhITZdcvTmleoBPoXaQ6uxqWpzQhJn2y_1GyVcD9cObGhRW1wlQ8zb0QIQiEUcmMfZ7LmFRjG33-EsVu6lW2boc1Oec_XIze57Yzuiy1mR_l1aqbHAbji4Q-PLhY4uBtNq6ZHQY2vTafg2t2VbmxczmHw10YXKHwqYC7bmS68i_yEO2x0iOr1XK-uRqsXWOLJGgv4UIrm1SmJX7IVC0WRu-OC_AfRUWHzELdY27CEaE9tMHkRkPJOETY3nMvDRxE8VKpQ1AW2aRFGr2GvL_S1m3AFXeCPl1XI9KyrmXWOfpPTIzCLw",
			"RS256",
			"",
			"https://login.microsoftonline.com/common/discovery/keys",
			goutils.ErrTokenMalformed,
		},
		{
			"ErrTokenSigningMethod",
			"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE1ODQyMDYyNTIsImV4cCI6MTg5OTczOTA1MiwiYXVkIjoid3d3LmV4YW1wbGUuY29tIiwic3ViIjoianJvY2tldEBleGFtcGxlLmNvbSIsIkdpdmVuTmFtZSI6IkpvaG5ueSIsIlN1cm5hbWUiOiJSb2NrZXQiLCJFbWFpbCI6Impyb2NrZXRAZXhhbXBsZS5jb20iLCJSb2xlIjpbIk1hbmFnZXIiLCJQcm9qZWN0IEFkbWluaXN0cmF0b3IiXX0.4dD8t29CgiL8adNgcAsSscdzx4EIp-rAxfGvbVM1xfk",
			"RS256",
			"",
			"",
			goutils.ErrTokenUnexpectedSigningMethod,
		},
		{
			"OKTokenHS256WithSignKey",
			"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE1ODQyNjQ0ODksImV4cCI6MjUzMDk0OTI4OSwiYXVkIjoid3d3LmV4YW1wbGUuY29tIiwic3ViIjoianJvY2tldEBleGFtcGxlLmNvbSIsIkdpdmVuTmFtZSI6IkpvaG5ueSIsIlN1cm5hbWUiOiJSb2NrZXQiLCJFbWFpbCI6Impyb2NrZXRAZXhhbXBsZS5jb20iLCJSb2xlIjpbIk1hbmFnZXIiLCJQcm9qZWN0IEFkbWluaXN0cmF0b3IiXX0.emaDzMycrC3xnMw3aZ4MRPtHiG10dfLqfNIL-1ax4Ww",
			"HS256",
			"qwertyuiopasdfghjklzxcvbnm123456",
			"",
			map[string]interface{}{
				"iss":       "Online JWT Builder",
				"iat":       1584264489,
				"exp":       2530949289,
				"aud":       "www.example.com",
				"sub":       "jrocket@example.com",
				"GivenName": "Johnny",
				"Surname":   "Rocket",
				"Email":     "jrocket@example.com",
				"Role": []string{
					"Manager",
					"Project Administrator",
				},
			},
		},
	}

	jwtToken := goutils.NewJWTToken()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			jwtToken.SetSignMethod(tc.signMethod)
			jwtToken.SetSignKey(tc.signKey)
			jwtToken.SetJwksURI(tc.jwksURI)
			_, err := jwtToken.ParseToken(tc.token)

			if strings.HasPrefix(tc.name, "Err") {
				assert.Equal(t, err, tc.result)
				return
			}

			assert.Nil(t, err)
		})
	}
}
