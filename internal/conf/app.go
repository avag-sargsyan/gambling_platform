package conf

type App struct {
	GRPcAddress       string `env:"GP_GRPC_ADDRESS" envDefault:"0.0.0.0:50051"`
	GRPcAddressClient string `env:"GP_GRPC_ADDRESS_CLIENT" envDefault:"app:50051"`
	HTTPAddress       string `env:"GP_HTTP_ADDRESS" envDefault:":8080"`
	RedisAddress      string `env:"GP_REDIS_ADDRESS" envDefault:"redis:6379"`
}
