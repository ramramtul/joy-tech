package schedule

import (
	"encoding/json"
	"joy-tech/models"
	"net/http"
	"strings"
)

func HandleSchedulePickUp(w http.ResponseWriter, r *http.Request) {
	var request models.ScheduleRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(strings.TrimSpace(request.EditionNumber)) == 0 || request.UserID == 0 || request.PickUpSchedule.IsZero() {
		http.Error(w, "Fields cannot be empty", http.StatusUnprocessableEntity)
		return
	}

	scheduled, err := SchedulePickUpBook(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(scheduled)
}
