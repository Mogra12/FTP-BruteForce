package modules

import "flag"

type Config struct {
	MaxConcurrent int
	SleepDuration int
	WordlistPath  string
	Hostname      string
	TlsFlag       bool
}

func LoadConfig() *Config {
	var config Config
	flag.IntVar(&config.MaxConcurrent, "maxConcurrent", 5, "Maximum number of concurrent connections")
	flag.BoolVar(&config.TlsFlag, "tls", false, "If true, uses FTPS (FTP over TLS)")
	flag.StringVar(&config.WordlistPath, "wl", "", "Path to the wordlist file")
	flag.StringVar(&config.Hostname, "h", "", "FTP server address (e.g., hostname:21)")
	flag.IntVar(&config.SleepDuration, "time", 1, "Number of seconds to sleep between login attempts when lwr is true")
	flag.Parse()
	
	return &config
}