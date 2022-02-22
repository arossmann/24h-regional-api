package store

import (
	"github.com/arossmann/24h-regional-api/db"
	"github.com/gofiber/fiber/v2"
)

func GetStores(c *fiber.Ctx) error {
	var loadedStores, _ = db.GetAllStores()
	return c.JSON(loadedStores)
}

func GetStore(c *fiber.Ctx) error {
	var loadedStore, _ = db.GetStoreByID(c.Params("id"))
	return c.JSON(loadedStore)
}

func NewStore(c *fiber.Ctx) error {
	return db.Create(c)
}

func DeleteStore(c *fiber.Ctx) error {
	return db.Delete(c)
}

func UpdateStore(c *fiber.Ctx) error {
	return db.Update(c)
}
