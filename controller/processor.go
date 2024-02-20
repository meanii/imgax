package controller

import (
	"github.com/gofiber/fiber/v2"

	"github.com/meanii/imgax/engine"
)

func Processor(c *fiber.Ctx) error {
	rembgEngine := new(engine.RembgEngine)
	if err := c.QueryParser(rembgEngine); err != nil {
		return err
	}
	savedPath, err := rembgEngine.Do()
	if err != nil {
		return err
	}
	return c.SendFile(savedPath)
}
