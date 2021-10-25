module lucasfrct.com/server

go 1.16

replace lucasfrct.com/login => ./src/login

replace lucasfrct.com/account => ./src/account

replace lucasfrct.com/server/router/router => ./src/server/router

replace lucasfrct.com/server/server => ./src/server

require github.com/googollee/go-socket.io v1.6.1 // indirect
