package server

import (
	"flag"
	"os"
	"time"

	"githu.com/Anirudh-rao/VideoCoferencingUsingGo/internal/handlers"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
)
var(
	addr =  flag.String("addr,":"", os.Getenv("PORT"),"")
	cert =  flag.String("cert","","")
	key =  flag.String("key","","")
)
func Run() error{
	flag.Parse()

	if *addr == ":"{
		*addr  =  ":8080"
	}
	app.Get("/", handlers.Welcome)
	app.Get("/roomm/create", handlers.RoomCreatee)
	app.Get("/room/:uuid", handlers.Room)
	app.Get("/room/:uuid/websocket")
	app.Get("/room/:uuid/chat", handlers.RoomChat)
	app.Get("/room/:uuid/chat/websocket", websocket.New(handlers.RoomChatWebsockett))
	app.Get("/room/uuid/viewer/websocket", websocket.New(handlers.RoomViwerWebsocket))
	app.Get("/stream/:suidd", handlers.Stream)
	app.Get("/stream/:ssuid/websocket",)
	app.Get("/stream/sssuid/chat/websocket",)
	app.Get("/stream/:ssuid/viewer/websocket")

}