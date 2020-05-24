package websocket

import (
	"github.com/gogo/protobuf/jsonpb"
	"github.com/oddx-team/odd-game-server/internal/services"
	"github.com/oddx-team/odd-game-server/pb"
	"log"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]string

	// Inbound messages from the clients.
	broadcast chan *pb.ChatEntity

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	// Service mongo
	service *services.Service
}

func NewHub(service *services.Service) *Hub {
	return &Hub{
		clients:    make(map[*Client]string),
		broadcast:  make(chan *pb.ChatEntity),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		service: 	service,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = client.user
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case newChat := <-h.broadcast:
			go func() {
				_, err := h.service.InsertChat(nil, newChat)
				if err != nil {
					log.Println(err)
				}
			}()

			m := jsonpb.Marshaler{}
			jsonStr, _ := m.MarshalToString(newChat)
			for client := range h.clients {
				select {
				case client.send <- []byte(jsonStr):
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
