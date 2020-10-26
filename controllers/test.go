package controllers

import (
	"fiber/models"
	"fiber/models/redis"
	"fmt"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

// example
func Test(c *fiber.Ctx) error {
	return nil
}

// @doc https://gorm.io/docs/
func TestOrm(c *fiber.Ctx) error {
	users, rowsCount := models.GetUserByUid(47)
	if rowsCount == 0 {
		return c.JSON(Response{
			Status:  false,
			Message: "Failed",
			Data:    nil,
			Error:   Error{
				ErrorCode:    1,
				ErrorMessage: "User not found",
			},
		})
	}

	return c.JSON(Response{
		Status:  true,
		Message: "Success",
		Data:    users,
	})
}

// @doc https://github.com/go-redis/redis
func TestRedis(c *fiber.Ctx) error {
	r := redis.ConnectDb(0)
	//r.Set("test_key", 1, 10000000000) // cache redis 10s - expiration tinh bang nano second = 1/tỷ second
	d, err := r.Get("test_key").Result()
	if err != nil {
		return c.JSON(Response{
			Status:  false,
			Message: "Failed",
			Data:    nil,
			Error:   Error{
				ErrorCode:    1,
				ErrorMessage: "Key not found",
			},
		})
	}
	return c.JSON(Response{
		Status:  true,
		Message: "Success",
		Data:    d,
		Error:   Error{},
	})
}

func TestFastHttp(c *fiber.Ctx) error {
	//var body []byte // This buffer could be acquired from a custom buffer pool
	//
	//c := &fasthttp.HostClient{
	//	IsTLS: true,
	//	Addr:  "localhost:443",
	//	TLSConfig: &tls.Config{
	//		InsecureSkipVerify: true,
	//	},
	//}
	//
	//statusCode, body, err := c.Get(nil, "https://localhost:443/v1/users/get-one")
	//if err != nil {
	//	log.Fatalf("Error: %s", err)
	//}
	//if statusCode != fasthttp.StatusOK {
	//	log.Fatalf("Unexpected status code: %d. Expecting %d", statusCode, fasthttp.StatusOK)
	//}

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
	})
}

func TestMongoDb(c *fiber.Ctx) error {
	return nil
}