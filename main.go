package main

import "github.com/gin-gonic/gin"
import "github.com/jinzhu/gorm"
import _ "github.com/mattn/go-sqlite3"
import "strconv"

type Drink struct {
	gorm.Model
  Name   string
	Price int
	Amount int
}

// DB初期化
func dbInit() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
  if err != nil {
		panic("can't connect DB（fail dbInit）")
	}
  db.AutoMigrate(&Drink{})
  defer db.Close()
}

// DB index
func dbGetAll() []Drink {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("can't connect DB (dbGetAll())")
    }
    var drinks []Drink
    db.Order("created_at desc").Find(&drinks)
    db.Close()
    return drinks
}

// DB show
func dbGetOne(id int) Drink {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("can't connect DB (dbGetOne())")
    }
    var drink Drink
    db.First(&drink, id)
    db.Close()
    return drink
}

// DB update
func dbUpdate(id int, amount int) {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("can't connect DB（dbUpdate)")
    }
    var drink Drink
    db.First(&drink, id)
		drink.Amount = amount
    db.Save(&drink)
    db.Close()
}

func main() {
	r := gin.Default()

	r.GET("/list", func(c *gin.Context) {
		drinks := dbGetAll()
		c.JSON(200, gin.H{
			"drinks": drinks,
		})
	})

	r.PUT("/buy/:id", func(c *gin.Context) {
			n := c.Param("id")
			id, err := strconv.Atoi(n)
			if err != nil {
					panic(err)
			}
			amount := dbGetOne(id).Amount -1
			dbUpdate(id, amount)

			c.JSON(200, gin.H{
				"drinks": dbGetOne(id),
			})
	})

	r.Run()
}
