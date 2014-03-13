package main

import (
	"goepoll"
)

func main() {
	goepoll.RunEpoll("127.0.0.1:8088")
}
