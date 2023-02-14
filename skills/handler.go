package skills

import (
	"encoding/json"
	"net/http"
	"time"

	"portfolio.com/db"
	"portfolio.com/model"
)

// ListSkills returns the list of skills.
func ListSkills(w http.ResponseWriter, r *http.Request) {
	var skill model.Skill
	var rsp model.Response
	var skills []model.Skill

	sqlDB := db.Connect()
	defer sqlDB.Close()

	rows, err := sqlDB.Query("SELECT id, name, years, rating FROM skills")
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		err = rows.Scan(&skill.ID, &skill.Name, &skill.Years, &skill.Rating)
		if err != nil {
			panic(err.Error())
		}

		skills = append(skills, skill)
	}

	rsp.Status = 200
	rsp.Message = "Success"
	rsp.Data = skills

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(rsp)
}

// AddSkill inserts a new skill.
func AddSkill(w http.ResponseWriter, r *http.Request) {
	var rsp model.Response

	db := db.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err.Error())
	}
	name := r.FormValue("name")
	years := r.FormValue("years")
	rating := r.FormValue("rating")
	updated := time.Now()

	_, err = db.Exec(
		"INSERT INTO skills(name, years, rating, updated) VALUES(?, ?, ?, ?)",
		name, years, rating, updated,
	)
	if err != nil {
		panic(err.Error())
	}

	rsp.Status = 200
	rsp.Message = "Success."

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(rsp)
}
