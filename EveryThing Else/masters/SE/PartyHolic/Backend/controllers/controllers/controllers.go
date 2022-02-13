package controllers

import (
	"database/database"
	"database/sql"
	"fmt"
	"geo/geo"
	"models/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAddresses(c *gin.Context) {
	var address []models.Addresses
	database.DB.Raw("Select * from addresses").Scan(&address)
	c.JSON(http.StatusOK, gin.H{"message": &address})
}

func GetAddressById(c *gin.Context) {
	var address models.Addresses

	if err := database.DB.Where("address_id = ?", c.Param("address_id")).First(&address).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": address})
}

func AddAddress(c *gin.Context) {
	var input models.Addresses
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	address_string := string(input.Lane_apt + "," + input.City + ", " + input.State + ", " + input.Country)
	location := (geo.GeoAddress(address_string))

	address := models.Addresses{Lane_apt: input.Lane_apt,
		City:      input.City,
		State:     input.State,
		Country:   input.Country,
		Latitude:  location[0],
		Longitude: location[1],
	}
	database.DB.Create(&address)

	c.JSON(http.StatusOK, gin.H{"message": input})
}

func AddUser(c *gin.Context) {

	var input models.Users
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.Users{User_id: input.User_id,
		First_name: input.First_name,
		Last_name:  input.Last_name,
		Address_id: input.Address_id,
		Gender:     input.Gender,
		Bio:        input.Bio,
	}

	database.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"message2": user})

}

func AddParty(c *gin.Context) {
	var input models.Parties
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// new_address_id = //database fetch

	party := &models.Parties{

		Party_name: input.Party_name,
		Host_id:    input.Host_id,
		Address_id: input.Address_id,

		Tags:        input.Tags,
		Description: input.Description,

		Start_time: input.Start_time,
		End_time:   input.End_time,

		Image_id:       input.Image_id,
		Attendee_count: input.Attendee_count,

		Latitude:  input.Latitude,
		Longitude: input.Longitude,
	}

	database.DB.Create(&party)

	c.JSON(http.StatusOK, gin.H{"message": party})

}

func GetParties(c *gin.Context) {
	var request models.PartiesRequest
	var parties []models.Party

	c.BindJSON(&request)

	distance_calculation := "(((acos(sin((?*3.414/180)) * sin((p.latitude*3.414/180))+cos((?*3.414/180))*cos((p.latitude*3.414/180))*cos(((?-p.longitude)*3.414/180))))*180/3.414)*60*1.1515*1609.344)/1609.34"
	distance_string := "select p.party_id, p.Party_name, concat_ws(' ', u.first_name, u.last_name) as Host_name, p.attendee_count, round(" + distance_calculation + ",2) as Distance, p.image_id from parties p join users u on p.host_id = u.user_id order by 4 limit 20"

	database.DB.Raw(distance_string, request.Location.Latitude, request.Location.Latitude, request.Location.Longitude).Scan(&parties)
	c.JSON(http.StatusOK, gin.H{"parties": parties})

}

func GetParty(c *gin.Context) {
	var party_details models.FullPartyDetails

	party_columns := "parties.Party_id, parties.party_name, parties.Start_time, parties.end_time, parties.tags, parties.description, parties.image_id, parties.attendee_count as interested_people,"
	user_columns := "users.first_name, users.last_name,"
	address_columns := "addresses.Lane_apt, addresses.City, addresses.State, addresses.Country, addresses.Latitude, addresses.Longitude"

	user_join := "JOIN users on users.user_id = parties.host_id"
	address_join := "Join addresses on addresses.address_id = parties.address_id"

	database.DB.Table("parties").Select(party_columns+user_columns+address_columns).Where("parties.party_id = ?", c.Param("party_id")).Joins(user_join).Joins(address_join).Find(&party_details)

	c.JSON(http.StatusOK, gin.H{"data": party_details})

}

func CancelParty(c *gin.Context) {
	id := c.Param("party_id")

	sql_connection, err2 := sql.Open("mysql", "root:@tcp(0.0.0.0:3306)/partyholic")
	res, err := sql_connection.Query("insert into cancelled_parties (select * from parties where party_id=?)", id)
	defer sql_connection.Close()
	fmt.Println(res, err2, err)

	database.DB.Delete(&models.Parties{}, id)

}

func AttendParty(c *gin.Context) {
	var attendee models.AttendeeList
	c.BindJSON(&attendee)
	database.DB.Create(&attendee)

	var count int

	sql_connection, _ := sql.Open("mysql", "root:@tcp(0.0.0.0:3306)/partyholic")
	sql_connection.QueryRow("select count(party_id) from attendee_lists where party_id=?", attendee.Party_id).Scan(&count)
	fmt.Println(count)
	sql_connection.Query("update parties set attendee_count = ? where party_id =?", count+1, attendee.Party_id)
	defer sql_connection.Close()

}

func CancelAttendance(c *gin.Context) {
	var attendee models.AttendeeList
	c.BindJSON(&attendee)

	database.DB.Where("user_id = ? and party_id = ?", attendee.User_id, attendee.Party_id).Delete(models.AttendeeList{})

	var count int
	sql_connection, _ := sql.Open("mysql", "root:@tcp(0.0.0.0:3306)/partyholic")
	sql_connection.QueryRow("select count(party_id) from attendee_lists where party_id=?", attendee.Party_id).Scan(&count)
	fmt.Println(count)
	sql_connection.Query("update parties set attendee_count = ? where party_id =?", count, attendee.Party_id)
	defer sql_connection.Close()
	c.JSON(http.StatusOK, gin.H{"data": attendee})

}

// func options(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{"message": "options Called"})
// }
