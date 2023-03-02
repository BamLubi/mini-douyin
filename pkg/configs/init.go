package configs

func InitConfigs(logpath string, loglevel string) {
	InitLoggerOnly(logpath, loglevel)
	initDB()
	initRedis()

}

func InitLoggerOnly(path string, level string) {
	initLogger(path, level)
}
