package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/golang-jwt/jwt/v5"
)

// ValidateJWTMiddleware extracting the request headers
// extracting our claims
// validating everything
func ValidateJWTMiddleware(next func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)) func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		// extract the headers
		token := extractTokenFromHeaders(request.Headers)

		if token == "" {
			return events.APIGatewayProxyResponse{
					Body:       "Missing Authorization token",
					StatusCode: http.StatusUnauthorized,
				},
				fmt.Errorf("invalid token")

		}

		// we need to parse our token for its claims
		claims, err := parseToken(token)
		if err != nil {
			return events.APIGatewayProxyResponse{
					Body:       "Unauthorized user",
					StatusCode: http.StatusUnauthorized,
				},
				fmt.Errorf("unauthorized user - %w", err)
		}

		expires := int64(claims["expires"].(float64))
		if time.Now().Unix() > expires {
			return events.APIGatewayProxyResponse{
					Body:       "Token is expired",
					StatusCode: http.StatusUnauthorized,
				},
				fmt.Errorf("token is expired")
		}
		return next(request)

	}
}

func extractTokenFromHeaders(headers map[string]string) string {
	authHeader, ok := headers["Authorization"]
	if !ok {
		return ""
	}

	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		return ""
	}
	return strings.TrimSpace(splitToken[1])
}

func parseToken(tokenString string) (jwt.MapClaims, error) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		secret := "secret"
		return []byte(secret), nil
	})

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}
