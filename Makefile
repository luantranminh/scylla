#!make
include .env
export $(shell sed 's/=.*//' .env)

dev:
	echo ${DB_USERNAME}
	bee run