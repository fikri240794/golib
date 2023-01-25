package status_code

import (
	"net/http"
)

type StatusCode int64

const (
	OK                    StatusCode = 22000
	CREATED               StatusCode = 22011
	BAD_REQUEST           StatusCode = 44000
	UNAUTHORIZED          StatusCode = 44011
	FORBIDDEN             StatusCode = 44033
	NOT_FOUND             StatusCode = 44044
	REQUEST_TIMEOUT       StatusCode = 44088
	TOO_MANY_REQUEST      StatusCode = 44299
	INTERNAL_SERVER_ERROR StatusCode = 55000
	BAD_GATEWAY           StatusCode = 55022
	GATEWAY_TIMEOUT       StatusCode = 55044
)

var statusCodes map[string]StatusCode = map[string]StatusCode{
	"ok":                    OK,
	"created":               CREATED,
	"bad request":           BAD_REQUEST,
	"unauthorized":          UNAUTHORIZED,
	"forbidden":             FORBIDDEN,
	"not found":             NOT_FOUND,
	"request timeout":       REQUEST_TIMEOUT,
	"too many request":      TOO_MANY_REQUEST,
	"internal server error": INTERNAL_SERVER_ERROR,
	"bad gateway":           BAD_GATEWAY,
	"gateway timeout":       GATEWAY_TIMEOUT,
}

func (code StatusCode) ToString() string {
	var s string

	for k, v := range statusCodes {
		if v == code {
			s = k
		}
	}

	return s
}

func (code StatusCode) ToHttpStatusCode() int {
	switch code {
	case CREATED:
		return http.StatusCreated
	case BAD_REQUEST:
		return http.StatusBadRequest
	case UNAUTHORIZED:
		return http.StatusUnauthorized
	case FORBIDDEN:
		return http.StatusForbidden
	case NOT_FOUND:
		return http.StatusNotFound
	case REQUEST_TIMEOUT:
		return http.StatusRequestTimeout
	case TOO_MANY_REQUEST:
		return http.StatusTooManyRequests
	case INTERNAL_SERVER_ERROR:
		return http.StatusInternalServerError
	case BAD_GATEWAY:
		return http.StatusBadGateway
	case GATEWAY_TIMEOUT:
		return http.StatusGatewayTimeout
	default:
		return http.StatusOK
	}
}
