package log

// init default logger
func init() {
	log, _ = NewZapLogger(Configuration{
		EnableConsole:     true,
		ConsoleJSONFormat: true,
		ConsoleLevel:      Info,
	})
}
