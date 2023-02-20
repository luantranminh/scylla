#!make
include .env
export $(shell sed 's/=.*//' .env)

dev:
	bee run