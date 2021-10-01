package todo

import (
  "fmt"
  "time"
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
)

type Todo struct {
  gorm.Model  
  Title string
  Done  bool
  Due   time.Time
}

func Init() *gorm.DB {
  db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
  if err != nil {
    panic("Database Error")
  }

  db.AutoMigrate(&Todo{})
  return db
}

// TODO Add better show format
func ShowAll() {
  db := Init()
  var todos []Todo

  db.Find(&todos)

  fmt.Println(todos)
}

func Add(title string) {
  db := Init()
  db.Create(&Todo{Title: title, Done: false, Due: time.Now()})
}

func Delete(id string) {
  fmt.Println("Delete", id)
  // db := Init()
}

func Done(id string) {
  fmt.Println("Done", id)
  // db := Init()
}
