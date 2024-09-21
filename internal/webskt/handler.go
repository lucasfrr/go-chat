package webskt

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	hub *Hub
}

func NewHandler(h *Hub) *Handler {
	return &Handler{
		hub: h,
	}
}

type CreateRoomRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var request CreateRoomRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.hub.Rooms[request.ID] = &Room{
		ID:      request.ID,
		Name:    request.Name,
		Clients: make(map[string]*Client),
	}

	c.JSON(http.StatusCreated, request)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) JoinRoom(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roomID := c.Param("roomId")
	clientID := c.Query("userId")
	username := c.Query("username")

	client := &Client{
		Conn:     conn,
		Message:  make(chan *Message, 10),
		ID:       clientID,
		RoomID:   roomID,
		Username: username,
	}

	msg := &Message{
		Content:  "A new user has been joined in the room",
		RoomID:   roomID,
		Username: username,
	}

	//Register a new client through the register channel
	h.hub.Register <- client

	// Broadcast that message
	h.hub.Broadcast <- msg

	go client.writeMessage()
	client.readMessage(h.hub)
}

type RoomResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) GetRooms(c *gin.Context) {
	rooms := make([]RoomResponse, 0)

	for _, room := range h.hub.Rooms {
		rooms = append(rooms, RoomResponse{
			ID:   room.ID,
			Name: room.Name,
		})
	}

	c.JSON(http.StatusOK, rooms)
}

type ClientsResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (h *Handler) GetRoomClients(c *gin.Context) {
	var clients []ClientsResponse
	roomId := c.Param("roomId")

	if _, ok := h.hub.Rooms[roomId]; !ok {
		clients = make([]ClientsResponse, 0)
		c.JSON(http.StatusOK, clients)
	}

	for _, client := range h.hub.Rooms[roomId].Clients {
		clients = append(clients, ClientsResponse{
			ID:       client.ID,
			Username: client.Username,
		})
	}

	c.JSON(http.StatusOK, clients)
}
