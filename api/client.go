// api/client.go
package api

import (
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: 10 * time.Second, // 🟢 Timeout for HTTP requests
}

func GetClient() *http.Client {
	return client
} /*GetClient()*/
