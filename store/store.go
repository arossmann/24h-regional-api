package store

import (
	"github.com/arossmann/24h-regional-api/db"
	"github.com/arossmann/24h-regional-api/entity"
	"github.com/gofiber/fiber/v2"
)

func GetStores(c *fiber.Ctx) error {
	var loadedStores, _ = db.GetAllStores()
	return c.JSON(loadedStores)
}

func GetStore(c *fiber.Ctx) error {
	var store entity.Store
	var loadedStore, _ = db.GetStoreByID(store.ID)
	return c.JSON(loadedStore)
}

func NewStore(c *fiber.Ctx) error {
	var store entity.Store
	id, _ := db.Create(&store)
	return c.JSON(id)
}

func DeleteStore(c *fiber.Ctx) error {
	var store entity.Store
	savedStore, _ := db.Update(&store)
	return c.JSON(savedStore)
}
