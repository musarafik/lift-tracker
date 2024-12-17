package repo

import (
	"sync"
)

var repoLocks []*sync.Mutex

type Exercise struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Sets int    `json:"sets"`
	Reps int    `json:"reps"`
}

type ExercisesRepo struct {
	exercises map[int]Exercise
	nextId    int
	lock      *sync.Mutex
}

func InitExercisesRepo() *ExercisesRepo {
	var lock *sync.Mutex = new(sync.Mutex)
	repoLocks = append(repoLocks, lock)

	return &ExercisesRepo{
		exercises: make(map[int]Exercise),
		nextId:    1,
		lock:      lock,
	}
}

func (r *ExercisesRepo) GetAllExercises() []Exercise {
	allExercises := make([]Exercise, 0, len(r.exercises))

	r.lock.Lock()
	defer r.lock.Unlock()

	for _, exercise := range r.exercises {
		allExercises = append(allExercises, exercise)
	}

	return allExercises
}

func (r *ExercisesRepo) AddExercise(name string, reps int, sets int) {
	exercise := Exercise{
		ID:   r.nextId,
		Name: name,
		Sets: sets,
		Reps: reps,
	}

	r.nextId = r.nextId + 1

	r.lock.Lock()
	defer r.lock.Unlock()

	r.exercises[exercise.ID] = exercise
}
