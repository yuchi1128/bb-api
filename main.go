package main

import (
	"fmt"
	// "todo_app/config"
	"todo_app/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DBName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.Password = "testtest"
	// fmt.Println(u)

	// u.CreateUser()

	// u, _ := models.GetUser(1)
	// fmt.Println(u)

	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// user, _ := models.GetUser(2)

	// user.CreateTodo("First Todo")

	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	user, _ := models.GetUser(2)
	user.CreateTodo("Second Todo")

	todos, _ := user.GetTodos()
	for _, v := range todos {
		fmt.Println(v)
	}

}
