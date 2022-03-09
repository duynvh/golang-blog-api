package skuser

import (
	"golang-blog-api/common"
	"golang-blog-api/component"
	log "golang-blog-api/log"

	socketio "github.com/googollee/go-socket.io"
)

type LocaltionData struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func OnUserUpdateLocation(appCtx component.AppContext, requester common.Requester) func(s socketio.Conn, localtion LocaltionData) {
	return func(s socketio.Conn, location LocaltionData) {
		log.Print("User update localtion: user id is ", requester.GetUserId(), " at location", location)
	}
}
