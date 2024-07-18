module main

replace api => ./api

replace server => ./server

go 1.22.5

require (
	api v0.0.0-00010101000000-000000000000
	server v0.0.0-00010101000000-000000000000
)
