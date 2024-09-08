package main

import "gin_test/webook/internal/web"

func main() {
	server := web.RegisterRouters()

	_ = server.Run(":8080")
}
