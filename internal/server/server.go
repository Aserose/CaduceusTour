package server

import (
	"github.com/Aserose/CaduceusTour/internal/middleware"
	"github.com/Aserose/CaduceusTour/pkg/logger"
	"net/http"
)

type Server interface {
	Run(port string)
}

type webhookReqBody struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

type server struct {
	server     *http.Server
	middleware middleware.Middleware
	reqBody    webhookReqBody
	log        logger.Logger
}

func NewServer(handler middleware.Middleware, log logger.Logger) Server {
	return &server{
		middleware: handler,
		log:        log}
}
