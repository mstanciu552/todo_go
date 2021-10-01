package todo

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
  "strconv"
  "strings"
)

// TODO Change time date format to only date
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

func (todo *Todo) format() {
  y, m, d := todo.Due.Date()
  fmt.Printf("%d %s %d-%d-%d %t", todo.ID, todo.Title, y, m, d, todo.Done)
}

func ShowAll() {
	fmt.Println("ID\tTitle\tDue\tDone")
	db := Init()
	var todos []Todo
	db.Find(&todos)

	for _, todo := range todos {
		todo.format()
	}
}

// TODO Change time date format to only date
func Add(title string) {
	db := Init()
	db.Create(&Todo{Title: title, Done: false, Due: time.Now()})
}

func getInt(i int, e error) int { return i }

func Due(id string, date string) {
  db := Init()
  var date_split []string = strings.Split(date, "-")
  y, m, d := date_split[0], date_split[1], date_split[2]


  var dt time.Time = time.Date(getInt(strconv.Atoi(y)), time.Month(getInt(strconv.Atoi(m))), getInt(strconv.Atoi(d)), 0, 0, 0, 0, nil)
  var todo Todo
  db.First(&todo, id)
  db.Model(&todo).Update("Due", dt)
}

func Delete(id string) {
	db := Init()
	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo, id)
}

func Done(id string) {
	db := Init()
	var todo Todo
	db.First(&todo, id)
	db.Model(&todo).Update("Done", true)
}

func ClearDB() {
	db := Init()
	db.Migrator().DropTable(&Todo{})
	db.AutoMigrate(&Todo{})
}
