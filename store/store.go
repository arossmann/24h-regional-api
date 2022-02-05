package store

import (
	"github.com/arossmann/24h-regional-api/db"
	"github.com/arossmann/24h-regional-api/entity"
	"github.com/gofiber/fiber"
)

func GetStores(c *fiber.Ctx) {
	var loadedStores, _ = db.GetAllStores()
	c.JSON(loadedStores)
}

func GetStore(c *fiber.Ctx) {
	var store entity.Store
	var loadedStore, _ = db.GetStoreByID(store.ID)
	c.JSON(loadedStore)
}

func NewStore(c *fiber.Ctx) {
	var store entity.Store
	id, _ := db.Create(&store)
	c.JSON(id)
}

func DeleteStore(c *fiber.Ctx) {
	var store entity.Store
	savedStore, _ := db.Update(&store)
	c.JSON(savedStore)
}
