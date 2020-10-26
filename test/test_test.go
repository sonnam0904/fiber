package test

import (
	"crypto/tls"
	"github.com/valyala/fasthttp"
	"log"
	"testing"
)
// @doc https://gorm.io/docs/
func TestTestOrm(t *testing.T) {
	// test https with disable ssl
	c := &fasthttp.HostClient{
		IsTLS: true,
		Addr:  "localhost:443",
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	statusCode, _, err := c.Get(nil, "https://localhost:443/test/test-orm")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	if statusCode != fasthttp.StatusOK {
		log.Fatalf("Unexpected status code: %d. Expecting %d", statusCode, fasthttp.StatusOK)
	}
}

// @doc https://github.com/go-redis/redis
func BenchmarkTestOrm(b *testing.B) {
	for i:=0; i < b.N; i++ {
		c := &fasthttp.HostClient{
			IsTLS: true,
			Addr:  "localhost:443",
			TLSConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}

		statusCode, _, err := c.Get(nil, "https://localhost:443/test/test-orm")
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		if statusCode != fasthttp.StatusOK {
			log.Fatalf("Unexpected status code: %d. Expecting %d", statusCode, fasthttp.StatusOK)
		}
	}
}

func TestTestRedis(t *testing.T)  {
	c := &fasthttp.HostClient{
		IsTLS: true,
		Addr:  "localhost:443",
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	statusCode, _, err := c.Get(nil, "https://localhost:443/test/test-redis")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	if statusCode != fasthttp.StatusOK {
		log.Fatalf("Unexpected status code: %d. Expecting %d", statusCode, fasthttp.StatusOK)
	}
}

func BenchmarkTestRedis(b *testing.B) {
	for i:=0; i < b.N; i++ {
		c := &fasthttp.HostClient{
			IsTLS: true,
			Addr:  "localhost:443",
			TLSConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}

		statusCode, _, err := c.Get(nil, "https://localhost:443/test/test-redis")
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		if statusCode != fasthttp.StatusOK {
			log.Fatalf("Unexpected status code: %d. Expecting %d", statusCode, fasthttp.StatusOK)
		}
	}
}
