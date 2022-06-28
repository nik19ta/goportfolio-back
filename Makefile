# Makefile

env-prepare: # create .env-file for secrets
	cp -n .example.env  .env

image-prepare: # create folder "images" if not exists 
	mkdir -p images

build: # build server
	go build -o ./.bin/app ./cmd/api/main.go

start: # start server
	./.bin/app

dev: # build and start server
	go build -o ./.bin/app ./cmd/api/main.go
	./.bin/app
