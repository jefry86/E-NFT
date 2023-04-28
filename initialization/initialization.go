package initialization

func New() {
	initViper()
	initZap()
	initRedis()
	initGorm()
}
