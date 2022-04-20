# Makefile

env-prepare: # create .env-file for secrets
	cp -n .env.example .env

image-prepare: # create folder inages
	mkdir images

build: # build server
	go build -o ./.bin/app ./cmd/api/main.go

start: # start server
	./.bin/app

dev: # build and start server
	go build -o ./.bin/app ./cmd/api/main.go
	./.bin/app
