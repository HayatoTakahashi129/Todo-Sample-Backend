package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

var FiberConfig = fiber.Config{
	ReadTimeout:  5 * time.Second,
	WriteTimeout: 25 * time.Second,
	ErrorHandler: fiber.DefaultErrorHandler, // TODO: エラーハンドリングの処理として何か必要ないかは要確認。
}
