package redis

// Config chỉ dùng cho redis package (không liên quan internal/config)
type Config struct {
	Host     string
	Port     int
	Password string
	DB       int
}
