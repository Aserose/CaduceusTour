package server

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func (s *server) Handler(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)

	if err := json.Unmarshal(body, &s.reqBody); err != nil {
		logrus.Errorf("server: decoding error, %v", err.Error())
	}

	menu := s.middleware.CreateMenu()

	s.middleware.Begin(s.reqBody.Message.Text, s.reqBody.Message.Chat.ID, menu)
	s.middleware.Request(s.reqBody.Message.Text, s.reqBody.Message.Chat.ID, menu)

}

func (s *server) Run(port string) {
	s.log.Info("server: start to listen")
	s.server = &http.Server{
		Addr:    port,
		Handler: http.HandlerFunc(s.Handler)}

	if err := s.server.ListenAndServe(); err != nil {
		s.log.Errorf("server: error %v", err.Error())
	}

}
