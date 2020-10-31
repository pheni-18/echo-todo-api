package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"

    "gorm.io/driver/mysql"
	"gorm.io/gorm"

    "fmt"
    "strconv"
)

type Todo struct {
    gorm.Model
    Title string `json:"title"`
    Description string `json:"description"`
    Done int `json:"done"`
}

func GetDBConn() *gorm.DB {
    dsn := "root:pw-secret@tcp(127.0.0.1:13306)/echo-todo-api?charset=utf8mb4&parseTime=True&loc=Local"
    DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        fmt.Println(err)
    }
    
    return DB
}

func main() {
    // Echo instance
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    db := GetDBConn()
    db.AutoMigrate(&Todo{})

    // Routes
    e.GET("/todo", getTodoList)
    e.POST("/todo", createTodo)
    e.GET("/todo/:id", getTodo)
    e.PATCH("/todo/:id", updateTodo)
    e.DELETE("/todo/:id", deleteTodo)

    // Start server
    e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func getTodoList(c echo.Context) error {
    var todoList []*Todo
    db := GetDBConn()
    db.Find(&todoList)

    return c.JSON(http.StatusOK, todoList)
}

func createTodo(c echo.Context) error {
    var todo Todo
    if err := c.Bind(&todo); err != nil {
        return err
    }

    db := GetDBConn()
    db.Create(&todo)

    return c.JSON(http.StatusCreated, todo)
}

func getTodo(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    var todo Todo
    db := GetDBConn()
    db.First(&todo, id)
    
    return c.JSON(http.StatusOK, todo)
}

func updateTodo(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))

    var afterTodo Todo
    if err := c.Bind(&afterTodo); err != nil {
        return err
    }
    
    var beforeTodo Todo
    db := GetDBConn()
    db.First(&beforeTodo, id)
    db.Model(&beforeTodo).Updates(&afterTodo)
    
    return c.JSON(http.StatusOK, afterTodo)
}

func deleteTodo(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    db := GetDBConn()
    db.Delete(&Todo{}, id)

    return c.NoContent(http.StatusNoContent)
}
