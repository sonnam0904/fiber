package main

import (
	"crypto/tls"
	"fiber/routes"
	"github.com/gofiber/fiber/v2"
	"log"
	"github.com/gofiber/template/html"
)

func main() {
	// register new engine
	engine := html.New("./views", ".html")

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")
	// => http://localhost:3000/static/hello.html
	// => http://localhost:3000/static/js/jquery.js
	// => http://localhost:3000/static/css/style.css

	// register routes
	routes.RegisterAPI(app)

	//go func() {
	//	//	_ = app.Listen(":3000/html/dist") // Start server with http://localhost:3000
	//	//}()

	// Create tls certificate
	//cer, err := tls.LoadX509KeyPair("certs/ssl.cert", "certs/ssl.key")
	//cer, err := tls.LoadX509KeyPair("certs/cert.pem", "certs/key.pem")
	cer, err := tls.LoadX509KeyPair("certs/ssl-cert-snakeoil.pem", "certs/ssl-cert-snakeoil.key")
	if err != nil {
		log.Fatal(err)
	}

	config := &tls.Config{Certificates: []tls.Certificate{cer}}

	// Create custom listener
	ln, err := tls.Listen("tcp", ":443", config)
	if err != nil {
		panic(err)
	}

	// Start server with https/ssl enabled on http://localhost:443
	log.Fatal(app.Listener(ln))
}