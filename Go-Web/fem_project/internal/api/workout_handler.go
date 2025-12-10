package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/adkham01/fem_project/internal/store"
	"github.com/go-chi/chi/v5"
)

type WorkoutHandler struct {
	workoutStore store.WorkoutStore
}

func NewWorkoutHandler(workoutStore store.WorkoutStore) *WorkoutHandler {
	return &WorkoutHandler{
		workoutStore: workoutStore,
	}
}

func (wh *WorkoutHandler) HandleGetWorkoutById(w http.ResponseWriter, r *http.Request) {
	paramsWorkoutId := chi.URLParam(r, "workoutId")
	if paramsWorkoutId == "" {
		http.NotFound(w, r)
		return
	}

	workoutId, err := strconv.ParseInt(paramsWorkoutId, 10, 64)
	if err != nil {
		http.NotFound(w, r)
	}

	fmt.Printf("this is workout id:%d\n", workoutId)
}

func (wh *WorkoutHandler) HandleCreateWorkout(w http.ResponseWriter, r *http.Request) {
	var workout store.Workout
	err := json.NewDecoder(r.Body).Decode(&workout)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		http.Error(w, "Failed to create workout", http.StatusInternalServerError)
		return
	}

	createdWorkout, err := wh.workoutStore.CreateWorkout(&workout)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		http.Error(w, "Failed to create workout", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdWorkout)
}
