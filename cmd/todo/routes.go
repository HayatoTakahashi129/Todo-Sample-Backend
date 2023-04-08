package main

import "github.com/gofiber/fiber/v2"

/**
* ルートの大元を設定するための関数。
* main.goはLambdaのハンドリングを扱っているので、純粋なFiberのレイヤーとして切り離す。
 */
func CreateRoute() *fiber.App {
	app := fiber.New(FiberConfig)
	// TODO: 暫定のルーティング設定を設置。後で置き換える。
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})
	return app
}
