package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-session/models"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connection = "mongodb+srv://golang:i0LTJWR3CdvM52yR@mycluster.qdloke1.mongodb.net/?retryWrites=true&w=majority"

const dbName = "golang_session"

const collectionName = "todos"

var tpl *template.Template

var collection *mongo.Collection

// this will run only one time when our application starts
func init() {
	tpl = template.Must(template.ParseFiles("index.html"))

	//client options
	clientOption := options.Client().ApplyURI(connection)

	// connect to mongo
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongoDb connection success")

	collection = (*mongo.Collection)(client.Database(dbName).Collection(collectionName))
}

func insertToDB(todo models.Todo) {
	inserted, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("data is inserted", inserted)
}

func getAll() []models.Todo {
	detail, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var todos []models.Todo
	// we have to decode or scan the details that we are fetching form somewhere else i.e. DB, r.body
	for detail.Next(context.Background()) {
		var todo models.Todo
		err := detail.Decode(&todo)
		if err != nil {
			panic(err)
		}
		todos = append(todos, todo)
	}
	defer detail.Close(context.Background())
	return todos
}

func HandlingTodosGETReq(w http.ResponseWriter, r *http.Request) {
	var Todos []models.Todo
	Todos = getAll()
	tpl.Execute(w, Todos)
}

func HandlingTodosPostReq(w http.ResponseWriter, r *http.Request) {
	var Todos []models.Todo
	var anotherTodo models.Todo
	// json.NewDecoder(r.Body).Decode(&anotherTodo)

	anotherTodo.Id, _ = strconv.Atoi(r.FormValue("id"))
	anotherTodo.TaskName = r.FormValue("taskname")

	anotherTodo.DateTime.Date = r.FormValue("date")
	anotherTodo.DateTime.Time = r.FormValue("time")

	go insertToDB(anotherTodo)
	Todos = getAll()
	Todos = append(Todos, anotherTodo)
	// json.NewEncoder(w).Encode(Todos)
	tpl.Execute(w, Todos)
}

func HandlingTodosDeleteReq(w http.ResponseWriter, r *http.Request) {
	var Todos []models.Todo
	query := r.URL.Query()

	// http://localhost:8080/todo?id=1&date=30-01-2023
	id, _ := strconv.Atoi(query.Get("id"))
	date := query.Get("date")
	taskName := query.Get("taskName")

	for index, todo := range Todos {
		if todo.Id == id && date == todo.DateTime.Date && taskName == todo.TaskName {
			Todos = append(Todos[:index], Todos[index+1:]...)
		}
	}

	json.NewEncoder(w).Encode(Todos)
}

func HandlingTodosPutReq(w http.ResponseWriter, r *http.Request) {
	var Todos []models.Todo

	if r.Method == http.MethodPut {
		query := r.URL.Query()
		id, _ := strconv.Atoi(query.Get("id"))

		var anothertodo models.Todo
		json.NewDecoder(r.Body).Decode(&anothertodo)

		for index, todo := range Todos {
			if todo.Id == id {
				Todos[index].DateTime.Date = anothertodo.DateTime.Date
				Todos[index].TaskName = anothertodo.TaskName
			}
		}
		json.NewEncoder(w).Encode(Todos)
	}
}

func HandlingHttpReq(w http.ResponseWriter, r *http.Request) {
	var Students []models.Student

	InitializingAndAddingToArray(&Students)

	if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(Students)
	}
	if r.Method == http.MethodPost {
		// creating anotherstudent structure
		var anotherStudent models.Student

		// decoding result from body
		json.NewDecoder(r.Body).Decode(&anotherStudent)
		// appending new student to the existing []struct
		Students = append(Students, anotherStudent)
		// sending the response
		json.NewEncoder(w).Encode(Students)
	}

	if r.Method == http.MethodDelete {
		// http://localhost:8080?id=2.. after this question mark ... it will pick as a query
		query := r.URL.Query()
		// we need to convert this string to int
		id, err := strconv.Atoi(query.Get("id"))
		if err != nil {
			fmt.Print("error occured", err)
			os.Exit(1)
		}

		for index, student := range Students {
			if id == student.Id {
				// splice and replace the students
				Students = append(Students[:index], Students[index+1:]...)
			}
		}
		json.NewEncoder(w).Encode(Students)
	}

	if r.Method == http.MethodPut {
		query := r.URL.Query()
		id, err := strconv.Atoi(query.Get("id"))
		if err != nil {
			fmt.Print("error occured", err)
			os.Exit(1)
		}

		var updateStudentDetail models.Student
		json.NewDecoder(r.Body).Decode(&updateStudentDetail)
		for index, student := range Students {
			if id == student.Id {
				Students[index].Name = updateStudentDetail.Name
				Students[index].Age = updateStudentDetail.Age
				Students[index].PhoneNo = updateStudentDetail.PhoneNo
			}
		}
		json.NewEncoder(w).Encode(Students)
	}
}
