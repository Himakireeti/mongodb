package main

import (
	"database/database"
	"geo/geo"

	"controllers/controllers"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {

	geo.GeoAddress("Gainesvile")

	database.Database_setup()
	r := gin.Default()
	r.Static("/images", "./images")
	r.Use(static.Serve("/", static.LocalFile("../Frontend/build", true)))

	r.GET("/api/addresses", controllers.GetAddresses)
	r.GET("/api/addresses/:address_id", controllers.GetAddressById)

	r.POST("/api/addAddress", controllers.AddAddress)

	r.POST("/api/newUser", controllers.AddUser)

	r.POST("/api/newParty", controllers.AddParty)
	r.POST("/api/parties", controllers.GetParties)
	r.GET("/api/getParty/:party_id", controllers.GetParty)
	r.GET("/api/cancelParty/:party_id", controllers.CancelParty)
	r.POST("/api/attendParty", controllers.AttendParty)
	r.POST("/api/cancelAttendance", controllers.CancelAttendance)

	// PORT environment variable was defined.
	r.Run()

}
