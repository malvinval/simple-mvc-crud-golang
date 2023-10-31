package config

import (
	"net/http"
)

func CreateServer(host string, port string) http.Server {
	return http.Server{
		Addr: host + ":" + port,
	}
}
