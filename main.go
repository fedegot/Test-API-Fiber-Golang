package main

import (
	"fmt"
	"fiber-mongo-api/configs"
	"fiber-mongo-api/routes"
	"github.com/gofiber/fiber/v2"
)

func main(){
app := fiber.New() //initialize Fiber application using new
configs.ConnectDB() //connecting to the DB
routes.UserRoute(app)// import of the Routes

app.Get("/", func(c *fiber.Ctx) error{
	return c.JSON(&fiber.Map{"data": "Using Fiber and MongoDB"}) //function to return Json file
})




fmt.Printf("Starting server at port 8080\n")
app.Listen(":8080");  //set Fiber to listen to port 8080
}

