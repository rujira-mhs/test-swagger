package main

import (
	"fmt"
	"net/http"

	//_ "swag-gin-demo/docs"
	//"./docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type book struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

var books = []book{
	{
		ID:     "1",
		Name:   "Harry Potter",
		Author: "J.K. Rowling",
		Price:  15.9,
	},
	{
		ID:     "2",
		Name:   "One Piece",
		Author: "Oda Eiichirō ii",
		Price:  2.99,
	},
	{
		ID:     "3",
		Name:   "Demon Slayer",
		Author: "Koyoharu Gotouge",
		Price:  2.99,
	},
}

// @Summary get all books
// @ID books
// @Produce json
// @Success 200 {object}
// @Router /books [get]

func getBooks(c *gin.Context) { //gin.context มีราบละเอียดของพวก request, validate
	c.JSON(http.StatusOK, books) //c.JSON ทำให้ struct เป็นรูปแบบของ JSON
	// http.statusOk คือเอาไว้ส่งกลับมาหา client ในที่นี้คือค่าคงที่ 200
	// books คือ Slice ของรายการหนังสือทั้งหมดที่จะตอบกลับไป

}

func getBookByID(c *gin.Context) {
	paramID := c.Param("id")
	for _, book := range books {
		if book.ID == paramID {
			c.JSON(http.StatusOK, book)
			fmt.Println("book.ID", book.ID, "paramID", paramID)
			return
		}
	}
	c.JSON(http.StatusNotFound, "data not found")
}

func addBook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil { //ถ้ามี err ให้รีเทิร์นกลับเลย
		return // เรียก BindJSON เพื่อผูก JSON ที่รับมากับ newBook
	}

	//เพิ่มรายการหนังสือเล่มใหม่เข้ามาใน slice
	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

// @Summary print hello
// @ID hello
// @Produce json
// @Success 200 {object}
// @Router /hello [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// @title Swagger Example API

func main() {
	r := gin.Default()

	r.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) //add swagger

	r.GET("/books", getBooks)
	r.GET("/books/:id", getBookByID)
	r.GET("/hello", Helloworld)
	r.POST("/books", addBook)

	r.Run("localhost:3000")
}
