package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Exercise struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Sets int    `json:"sets"`
	Reps int    `json:"reps"`
}

var exercisesRepo = make(map[int]Exercise)
var nextID = 1
var repoLock sync.Mutex

var allExercisesTemplate = template.Must(template.ParseFiles("templates/all_exercises.html"))
var createExerciseTemplate = template.Must(template.ParseFiles("templates/create_exercise.html"))

type AllExercisesData struct {
	AllExercises []Exercise
}

func handlePostExercises(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Unable to parse form")
		http.Error(w, err.Error(), http.StatusBadRequest) // TODO: need to create custom error page
		return
	}

	fmt.Println(r.Form)

	name := r.FormValue("name")

	sets, err := strconv.Atoi(r.FormValue("sets"))
	if err != nil {
		fmt.Println("Could not parse sets")
		http.Error(w, err.Error(), http.StatusBadRequest) // TODO: need to create custom error page
		return
	}

	reps, err := strconv.Atoi(r.FormValue("reps"))
	if err != nil {
		fmt.Println("Could not parse reps")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	exercise := Exercise{
		ID:   nextID,
		Name: name,
		Sets: sets,
		Reps: reps,
	}

	nextID = nextID + 1

	repoLock.Lock()
	defer repoLock.Unlock()

	exercisesRepo[nextID] = exercise
}

func allExercisesHandler(w http.ResponseWriter, r *http.Request) {
	repoLock.Lock()
	defer repoLock.Unlock()

	data := AllExercisesData{
		AllExercises: make([]Exercise, 0, len(exercisesRepo)),
	}

	for _, exercise := range exercisesRepo {
		data.AllExercises = append(data.AllExercises, exercise)
	}

	fmt.Println(data.AllExercises)

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

func exercisesApiHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// TODO
	case "POST":
		handlePostExercises(w, r)
	default:
		// TODO
	}

}

func main() {
	http.HandleFunc("/exercises/all", allExercisesHandler)
	http.HandleFunc("/exercises/create", createExerciseHandler)
	http.HandleFunc("/api/exercises", exercisesApiHandler)

	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
