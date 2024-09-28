// Command: go build -tags=k8s -o webook.
//go:build k8s

package config

var Config = config{
	DB: DBConfig{
		DSN: "localhost:30006",
	},
	Redis: RedisConfig{
		Addr: "localhost:31379",
	},
}
