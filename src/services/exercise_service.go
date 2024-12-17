package services

import (
	"lift_tracker/src/repo"
	"strconv"
)

var exerciseRepo *repo.ExercisesRepo

func InitExercisesService() {
	exerciseRepo = repo.InitExercisesRepo()
}

func GetAllExercises() []repo.Exercise {
	return exerciseRepo.GetAllExercises()
}

func CreateExercise(nameRaw string, repsRaw string, setsRaw string) error {
	name := nameRaw

	sets, err := strconv.Atoi(setsRaw)
	if err != nil {
		return err
	}

	reps, err := strconv.Atoi(repsRaw)
	if err != nil {
		return err
	}

	exerciseRepo.AddExercise(name, sets, reps)

	return nil
}
