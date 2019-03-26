package lib

import "time"

type Output struct {
	OS       string
	Username string
	Hostname string
	Uptime   time.Duration
	Freeram  float64
	Totalram float64
}
