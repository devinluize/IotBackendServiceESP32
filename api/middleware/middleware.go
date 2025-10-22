package middleware

import (
	configenv "IotBackend/api/config"
	"IotBackend/api/helper"
	"IotBackend/api/payloads/responses"
	"context"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

func RouterMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Credentials", "true")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		tokenString := request.Header.Get("Authorization")
		part := strings.Split(tokenString, " ")
		if len(part) > 1 {
			tokenString = part[1]
		} else {
			helper.ReturnStandarResponses(writer, false, "Token not found from Header", nil)
			return
		}
		claims := configenv.JWTClaim{}

		myToken, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return configenv.JWT_KEY, nil
		})
		if err != nil {
			helper.ReturnAPIResponses(writer, responses.ApiResponseError{
				Message: "token invalid or expired",
				Success: false,
				Err:     nil,
			})
			return
		}
		//if claims.UserName != "devin" {
		//	helper.ReturnAPIResponses(writer, Responses.ApiResponseError{
		//		Message: "note devin",
		//		Success: false,
		//		Err:    nil,
		//	})
		//}
		if claims.IsVIP {
			helper.ReturnAPIResponses(writer, responses.ApiResponseError{
				Message: "not VIP",
				Success: false,
				Err:     claims.IsVIP,
			})
			//return
		}
		//if claims.UserRole != 1 {
		//	exceptions.NewAuthorizationException(writer, request, &exceptions.BaseErrorResponse{
		//		StatusCode: http.StatusUnauthorized,
		//		Message:    "You are not authorized to this endpoint",
		//		Err:       nil,
		//		Err:        err,
		//	})
		//	return
		//}
		if !myToken.Valid {
			helper.ReturnAPIResponses(writer, responses.ApiResponseError{
				Message: "Token invalid",
				Success: false,
				Err:     nil,
			})
			return
		}
		ctx := context.WithValue(request.Context(), "user_credential", helper.UserContext{
			UserName: claims.UserName,
			UserId:   claims.UserId,
		})
		//contexts := context.WithValue(request.Context(), "userName", claims.UserName)
		request = request.WithContext(ctx)
		handler.ServeHTTP(writer, request)
	})
}
