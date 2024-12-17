package main

import (
	"fmt"
	"html/template"
	"lift_tracker/src/repo"
	"lift_tracker/src/services"
	"log"
	"net/http"
)

var allExercisesTemplate = template.Must(template.ParseFiles("src/templates/all_exercises.html"))
var createExerciseTemplate = template.Must(template.ParseFiles("src/templates/create_exercise.html"))
var createScheduleTemplate = template.Must(template.ParseFiles("src/templates/create_schedule.html"))
var indexTemplate = template.Must(template.ParseFiles("src/templates/index.html"))

type AllExercisesData struct {
	AllExercises []repo.Exercise
}

func handlePostExercises(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Unable to parse form")
		http.Error(w, err.Error(), http.StatusBadRequest) // TODO: need to create custom error page
		return
	}

	name := r.FormValue("name")
	sets := r.FormValue("sets")
	reps := r.FormValue("reps")

	err = services.CreateExercise(name, reps, sets)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	http.Redirect(w, r, "/exercises/all", http.StatusSeeOther)
}

func allExercisesHandler(w http.ResponseWriter, r *http.Request) {
	data := AllExercisesData{
		AllExercises: services.GetAllExercises(),
	}

	err := allExercisesTemplate.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}

func createExerciseHandler(w http.ResponseWriter, r *http.Request) {
	err := createExerciseTemplate.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func apiExerciseHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// TODO
	case "POST":
		handlePostExercises(w, r)
		// TODO
	default:
		// TODO
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := indexTemplate.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func handlePostSchedule(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func apiScheduleHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// TODO
	case "POST":
		handlePostSchedule(w, r)
		// TODO
	default:
		// TODO
	}
}

func createScheduleHandler(w http.ResponseWriter, r *http.Request) {
	err := createScheduleTemplate.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	services.InitExercisesService()
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/exercises/all", allExercisesHandler)
	http.HandleFunc("/exercises/create", createExerciseHandler)
	http.HandleFunc("/schedule/create", createScheduleHandler)
	http.HandleFunc("/api/exercises", apiExerciseHandler)
	http.HandleFunc("/api/schedule", apiScheduleHandler)

	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
