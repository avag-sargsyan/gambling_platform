package conf

type Log struct {
	Level string `env:"GP_LOGLEVEL" envDefault:"info"`
}
