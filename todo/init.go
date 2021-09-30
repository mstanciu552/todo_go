package todo

import (
  "fmt"
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
)

type Todo struct {
  gorm.Model  
  Title string
  Done  bool
  Due   string
}

func Init() {
  db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
  if err != nil {
    panic("Database Error")
  }

  db.AutoMigrate(&Todo{})

  // db.Create(&Todo{Title: "test2", Done: false, Due: "date"})

  var todo Todo
  db.First(&todo, "title = ?", "test2")
  fmt.Println(todo)
}

func ShowAll() {
  fmt.Println("Todo")
}

func Add() {
  fmt.Println("Add")
}

func Delete(id string) {
  fmt.Println("Delete", id)
}

func Done(id string) {
  fmt.Println("Done", id)
}
