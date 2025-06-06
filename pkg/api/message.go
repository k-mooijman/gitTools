package api

import "errors"

type Message struct {
	Name   string `json:"name"`
	Origin string `json:"origin"`
	Value  string `json:"value"`
}

func (s *Server) Sent(message Message) error {
	if s.WS == nil {
		return errors.New("Websocket not initialized")
	}
	return s.WS.WriteJSON(message)
}
