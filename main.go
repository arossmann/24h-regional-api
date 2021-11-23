package main

import (
	"github.com/arossmann/24h-regional-api/db"
	"github.com/arossmann/24h-regional-api/entity"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

// Get All Stores godoc
// @Summary returns the list of all stores
// @Description get the list of all stores.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/stores/ [get]
func handleGetStores(c *gin.Context) {
	/* var stores []Store
	var store Store
	store.Title = "Testprodukt"
	store.Description = "This is just a simple test store"
	stores = append(stores, store)
	c.JSON(http.StatusOK, gin.H{"stores": stores}) */

	var loadedStores, err = db.GetAllStores()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"stores": loadedStores})
}

func handleGetStore(c *gin.Context) {
	var store entity.Store
	if err := c.BindUri(&store); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	var loadedStore, err = db.GetStoreByID(store.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ID": store.ID, "Store": loadedStore})
}

func handleCreateStore(c *gin.Context) {
	var store entity.Store
	if err := c.ShouldBindJSON(&store); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	id, err := db.Create(&store)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func handleUpdateStore(c *gin.Context) {
	var store entity.Store
	if err := c.ShouldBindJSON(&store); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	savedStore, err := db.Update(&store)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"task:": savedStore})
}

func handleDeleteStore(c *gin.Context){
	var store entity.Store
	if err := c.ShouldBindJSON(&store); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	savedStore, err := db.Update(&store)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"task:": savedStore})
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/swagger/ [get]
func HealthGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "UP",
	})
}

// @title 24h Regional API
// @version 1.0
// @description API for 24h-regional.de
// @termsOfService http://swagger.io/terms/

// @contact.name Arne Rossmann
// @contact.url http://24h-regional.de
// @contact.email arne.rossmann@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8888
// @BasePath /api/v1
// @schemes http

func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		health := v1.Group("/health")
		{
			health.GET("", HealthGet)
		}
		/*location := v1.Group("/location")
		{
			location.GET(":address",handleGetGeoLocationFromAddress)
		}*/
		stores := v1.Group("/stores")
		{
			stores.GET(":id", handleGetStore)
			stores.GET("", handleGetStores)
			stores.POST("", handleCreateStore)
			stores.DELETE(":id", handleDeleteStore)
			stores.PUT(":id", handleUpdateStore)

		}
	}
	r.Run(":"+os.Getenv("PORT"))
}
