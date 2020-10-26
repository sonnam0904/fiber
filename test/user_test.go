package test

import (
	"crypto/tls"
	"github.com/valyala/fasthttp"
	"log"
	"testing"
)

func TestGetOne(t *testing.T) {
	// test https with disable ssl
	var body []byte // This buffer could be acquired from a custom buffer pool

	c := &fasthttp.HostClient{
		IsTLS: true,
		Addr:  "localhost:443",
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	statusCode, body, err := c.Get(nil, "https://localhost:443/v1/users/get-one")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	if statusCode != fasthttp.StatusOK {
		log.Fatalf("Unexpected status code: %d. Expecting %d", statusCode, fasthttp.StatusOK)
	}
	useResponseBody(body)
}

func useResponseBody(body []byte) {
	// Do something with body :)
}

func BenchmarkGetOne(b *testing.B) {
	for i:=0; i < b.N; i++ {
		// benchmark http
		//var body []byte // This buffer could be acquired from a custom buffer pool
		//// test call api by fasthttp
		//statusCode, body, err := fasthttp.Get(body, "http://localhost:3000/v1/users/get-one")
		//if err != nil {
		//	log.Fatalf("Error: %s", err)
		//}
		//if statusCode != fasthttp.StatusOK {
		//	log.Fatalf("StatusCode: %d", statusCode)
		//}

		// benchmark https
		var body []byte // This buffer could be acquired from a custom buffer pool

		c := &fasthttp.HostClient{
			IsTLS: true,
			Addr:  "localhost:443",
			TLSConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}

		statusCode, body, err := c.Get(nil, "https://localhost:443/v1/users/get-one")
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		if statusCode != fasthttp.StatusOK {
			log.Fatalf("Unexpected status code: %d. Expecting %d", statusCode, fasthttp.StatusOK)
		}
		useResponseBody(body)
	}
}