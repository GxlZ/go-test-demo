package main

import "google.golang.org/genproto/googleapis/cloud/redis/v1beta1"

func main() {

}

func NewUser(redisClient *redis.client) *user {
	return &user{}
}

type user struct {
}

func (this *user) Get() {

}

func (this *user) Set() {

}
