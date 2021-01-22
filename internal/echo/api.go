package echo

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/jxlwqq/go-restful/internal/auth"
	"github.com/jxlwqq/go-restful/pkg/log"

	"net/http"
)

func RegisterHandlers(r *mux.Router, logger *log.Logger, authMiddleware auth.Middleware) {
	res := resource{authMiddleware, logger}
	r.HandleFunc("/echo", res.echo)
}

type resource struct {
	m      auth.Middleware
	logger *log.Logger
}

var up websocket.Upgrader

func (res resource) echo(w http.ResponseWriter, r *http.Request) {
	conn, err := up.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	// 鉴权
	token := r.URL.Query().Get("token")
	claims, err := res.m.VerifyToken(token)
	if err != nil {
		return
	}
	id := claims.(jwt.MapClaims)["id"].(string)
	res.logger.Info("auth user:", id)

	// todo conn 与 id 绑定
	defer func() {
		conn.Close()
	}()

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			res.logger.Error("read:", err)
			break
		}
		res.logger.Info("recv: %v %s", mt, message)
		// do something
		err = conn.WriteMessage(mt, message)
		if err != nil {
			res.logger.Error("write:", err)
			break
		}
	}
}
