package service

// PingService returns the response for the ping endpoint.
func PingService() map[string]string {
    return map[string]string{
        "message": "pong",
    }
}