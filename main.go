package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleGetStores(c *gin.Context) {
	/* var stores []Store
	var store Store
	store.Title = "Testprodukt"
	store.Description = "This is just a simple test store"
	stores = append(stores, store)
	c.JSON(http.StatusOK, gin.H{"stores": stores}) */

	var loadedStores, err = GetAllStores()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"stores": loadedStores})
}

func handleGetStore(c *gin.Context) {
	var store Store
	if err := c.BindUri(&store); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	var loadedStore, err = GetStoreByID(store.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ID": store.ID, "Store": loadedStore})
}

func handleCreateStore(c *gin.Context) {
	var store Store
	if err := c.ShouldBindJSON(&store); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	id, err := Create(&store)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func handleUpdateStore(c *gin.Context) {
	var store Store
	if err := c.ShouldBindJSON(&store); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	savedStore, err := Update(&store)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"task:": savedStore})
}

func handleDeleteStore(c *gin.Context){
	var store Store
	if err := c.ShouldBindJSON(&store); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	savedStore, err := Update(&store)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"task:": savedStore})
}

func main() {
	r := gin.Default()
	r.GET("/stores/:id", handleGetStore)
	r.GET("/stores/", handleGetStores)
	r.PUT("/stores/", handleCreateStore)
	r.POST("/stores/", handleUpdateStore)
	r.Run()


	v1 := r.Group("/api/v1")
	{
		stores := v1.Group("/stores")
		{
			stores.GET(":id", handleGetStore)
			stores.GET("", handleGetStores)
			stores.POST("", handleCreateStore)
			stores.DELETE(":id", handleDeleteStore)
			stores.PATCH(":id", handleUpdateStore)

		}
	}
}
