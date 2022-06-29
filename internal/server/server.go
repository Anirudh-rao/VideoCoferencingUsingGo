package server

import (
	"flag"
	"os"

	"githu.com/Anirudh-rao/VideoCoferencingUsingGo/internal/handlers"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
)

var (
	addr =  flag.String("addr,":"", os.Getenv("PORT"),"")
	cert =  flag.String("cert","","")
	key =  flag.String("key","","")
)
func Run() error{
	flag.Parse()
	if *addr = ":"{
		*addr  = ":8080"
	}
	//Routes
	app.GET("/", handlers.Welcome)
	app.GET("/room/create", handlers.RoomCreate)
	app.GET("/room/:uuid", handlers.Room)
	app.GET("/room/:uuid/websocket", )
	app.GET("/room/:uuid/chat", handlers.RoomChat)
	app.GET("/room/:uuid/chat/websocket", websocket.New(handlers.RooomChatWeocket))
	app.GET("/room/:uuid/viewer/websocket", websocket.New(handlers.RoomViewerWebsocket))
	app.GET("/stream/:ssuid", handlers.Stream)
	app.GET("/stream/:ssuid/websocket", )
	app.GET("/stream/:ssuid/chat/websocket", )
	app.GET("/stream/:ssuid/viewer/websocket", )
}

func main(){

}