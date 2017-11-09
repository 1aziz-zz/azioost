package services

import "net/http"

type Route struct {
	Name, Method, Pattern string
	HandlerFunc           http.HandlerFunc
}
