package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type StudentAPI struct {
	ID string
	Name string
	NIM string
	Email string
	PhoneNumber string
	Class string
	Major string
	Faculty string
	University string
}

func main() {
	data := []StudentAPI{
		StudentAPI{
			ID:          "1",
			Name:        "sammi aldhi yanto",
			NIM:         "2003110000",
			Email:       "sammidev@mail.com",
			PhoneNumber: "12830790184",
			Class:       "A",
			Major:       "SI",
			Faculty:     "FMIPA",
			University:  "RIAU",
		},
		StudentAPI{
			ID:          "2",
			Name:        "aditya andika putra",
			NIM:         "20031123423",
			Email:       "adit@mail.com",
			PhoneNumber: "12830790184",
			Class:       "A",
			Major:       "SI",
			Faculty:     "FMIPA",
			University:  "RIAU",
		},
		StudentAPI{
			ID:          "3",
			Name:        "Ayatullah RJ",
			NIM:         "2003110000",
			Email:       "ayat@mail.com",
			PhoneNumber: "12830790184",
			Class:       "A",
			Major:       "SI",
			Faculty:     "FMIPA",
			University:  "RIAU",
		},
		StudentAPI{
			ID:          "4",
			Name:        "Gusnur",
			NIM:         "200314230000",
			Email:       "gusnur@mail.com",
			PhoneNumber: "12830790184",
			Class:       "A",
			Major:       "SI",
			Faculty:     "FMIPA",
			University:  "RIAU",
		},
		StudentAPI{
			ID:          "5",
			Name:        "Abdul Rauf",
			NIM:         "20012311230",
			Email:       "rauf@mail.com",
			PhoneNumber: "12830790184",
			Class:       "A",
			Major:       "SI",
			Faculty:     "FMIPA",
			University:  "RIAU",
		},
	}

	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "GET" {
			result, err := json.Marshal(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(result)
			return
		}
	})

	http.HandleFunc("/student", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "POST" {
			id := r.FormValue("id")
			var result []byte
			var err error

			for _, each := range data {
				if each.ID == id {
					result, err = json.Marshal(each)
					if err != nil {
						log.Println("tidak ada data")
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					w.Write(result)
					return
				}
			}
		}
	})

	log.Println("OK! RUN SMOOTHLY")
	http.ListenAndServe(":8000", nil)
}