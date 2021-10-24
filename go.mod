module lucasfrct.com/server

go 1.16

replace lucasfrct.com/router => ./src

replace lucasfrct.com/login => ./src/login

replace lucasfrct.com/account => ./src/account

require (
	lucasfrct.com/account v0.0.0-00010101000000-000000000000 // indirect
	lucasfrct.com/login v0.0.0-00010101000000-000000000000 // indirect
	lucasfrct.com/router v0.0.0-00010101000000-000000000000
)
