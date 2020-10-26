package controllers

import (
	"bytes"
	"fiber/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

func GetOne(c *fiber.Ctx) error {
	// Parse the multipart form:
	if form, err := c.MultipartForm(); err == nil {
		// => *multipart.Form

		if token := form.Value["token"]; len(token) > 0 {
			// Get key value:
			fmt.Println(token[0])
		}

		// Get all files from "documents" key:
		files := form.File["documents"]
		// => []*multipart.FileHeader

		// Loop through files:
		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
			// => "tutorial.pdf" 360641 "application/pdf"

			// Save the files to disk:
			//if err := c.SaveFile(file, fmt.Sprintf("./%s", file.Filename)); err != nil {
			//	return err
			//}
		}
	}

	// test get from models
	users, _ := models.GetUserByUid(47)

	var body []byte // This buffer could be acquired from a custom buffer pool
	// test call api by fasthttp
	statusCode, body, err := fasthttp.Get(body, "https://raw.githubusercontent.com/sonnam0904/json-place-holder/master/hotel.json")
	if err != nil {
		fmt.Printf("Client get failed: %s\n", err)
		return nil
	}
	if statusCode != fasthttp.StatusOK {
		fmt.Printf("Expected status code %d but got %d\n", fasthttp.StatusOK, statusCode)
		return nil
	}

	return c.JSON(fiber.Map{
		"success": true,
		"test" : "ok",
		"param":    []interface{}{
			fiber.Map{
				"name": []interface{}{
					fiber.Map{
						"name": jsoniter.Get(body, "data", 0, "name").ToString(),
						"link_detail": jsoniter.Get(body, "data", 0, "link_detail").ToString(),
					},
				},
				"link_detail": jsoniter.Get(body, "data", 0, "link_detail").ToString(),
			},
		},
		"user":    users,
	})
}

// example
func request(uri string) (statusCode int, body []byte, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI("https://raw.githubusercontent.com/sonnam0904/json-place-holder/master/hotel.json")
	// fasthttp does not automatically request a gzipped response.
	// We must explicitly ask for it.
	req.Header.Set("Accept-Encoding", "gzip")

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	// Perform the request
	err = fasthttp.Do(req, resp)
	if err != nil {
		fmt.Printf("Client get failed: %s\n", err)
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		fmt.Printf("Expected status code %d but got %d\n", fasthttp.StatusOK, resp.StatusCode())
	}

	// Verify the content type
	contentType := resp.Header.Peek("Content-Type")
	if bytes.Index(contentType, []byte("application/json")) != 0 {
		fmt.Printf("Expected content type application/json but got %s\n", contentType)
	}

	// Do we need to decompress the response?
	contentEncoding := resp.Header.Peek("Content-Encoding")
	if bytes.EqualFold(contentEncoding, []byte("gzip")) {
		fmt.Println("Unzipping...")
		body, _ = resp.BodyGunzip()
	} else {
		body = resp.Body()
	}

	fmt.Printf("Response body is: %s", body)
	return statusCode, body , err
}
