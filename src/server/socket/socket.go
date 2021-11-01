package socket

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
)

type Socket struct {
	server *socketio.Server
}

func (s *Socket) New() *socketio.Server {

	s.server = socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{CheckOrigin: s.allowOriginFunc},
			&websocket.Transport{CheckOrigin: s.allowOriginFunc},
		},
	})

	s.server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		s.Emit("connectedId", s.ID())
		log.Println("connected id: ", s.ID())
		return nil
	})

	s.server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	s.server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		s.Emit(s.ID())
		log.Println("closed ID: ", s.ID(), " - ", reason)
	})

	return s.server
}

func (s *Socket) Event(event string, clojure func(s socketio.Conn, data string)) {
	go s.server.OnEvent("/", event, clojure)
}

func (s *Socket) Run() {

	go func() {
		if err := s.server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	defer s.server.Close()

	http.Handle("/socket.io/", s.server)
	http.Handle("/", http.FileServer(http.Dir("./pages")))

	log.Println("Socket- localhost:8000...")
	http.ListenAndServe(":8000", nil)

}

// Easier to get running with CORS. Thanks for help @Vindexus and @erkie
func (s *Socket) allowOriginFunc(r *http.Request) bool {
	return true
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func main() {
	log.Println("==================================================")
	log.Println("Servertools - 0.0.1")
	log.Println("==================================================")

	socket := Socket{}
	socket.New()

	socket.Event("get:page", func(s socketio.Conn, data string) {
		log.Println("get:pages ID: ", s.ID(), " - ", data)
		log.Println()

		url, _ := url.Parse(data)
		// proxy := httputil.NewSingleHostReverseProxy(url)
		log.Println(url.Scheme)
		// response, err := http.Get(data)
		// check(err)
		// defer response.Body.Close()
		// body, _ := ioutil.ReadAll(response.Body)
		// fmt.Println(string(body))

		s.Emit("get:page", "http://localhost:8000/page")
		// s.Emit("get:page", "server->host")
	})

	socket.Run()

}

/*
				   port :8080
					server
				  /		   \
				 /			\
				/	         \
		 client 	  		   client
port :4200->8000            port :5500

*/
