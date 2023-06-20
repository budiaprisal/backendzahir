package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type Contact struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var contacts []Contact

func GetContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Filter by name
	nameFilter := strings.ToLower(r.FormValue("name"))

	filteredContacts := make([]Contact, 0)
	for _, contact := range contacts {
		if nameFilter != "" && !strings.Contains(strings.ToLower(contact.Name), nameFilter) {
			continue
		}
		filteredContacts = append(filteredContacts, contact)
	}

	// Sort
	sortBy := r.FormValue("sort") // Sort by field
	switch sortBy {
	case "name":
		sort.SliceStable(filteredContacts, func(i, j int) bool {
			return strings.ToLower(filteredContacts[i].Name) < strings.ToLower(filteredContacts[j].Name)
		})
	case "created_at":
		sort.SliceStable(filteredContacts, func(i, j int) bool {
			return filteredContacts[i].CreatedAt.Before(filteredContacts[j].CreatedAt)
		})
	default:
		// No sorting specified, use default sorting by ID
		sort.SliceStable(filteredContacts, func(i, j int) bool {
			return filteredContacts[i].ID < filteredContacts[j].ID
		})
	}

	// Pagination
	page, _ := strconv.Atoi(r.FormValue("page"))
	pageSize, _ := strconv.Atoi(r.FormValue("page_size"))
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	startIndex := (page - 1) * pageSize
	if startIndex >= len(filteredContacts) {
		json.NewEncoder(w).Encode([]Contact{})
		return
	}

	endIndex := startIndex + pageSize
	if endIndex > len(filteredContacts) {
		endIndex = len(filteredContacts)
	}

	pagedContacts := filteredContacts[startIndex:endIndex]
	json.NewEncoder(w).Encode(pagedContacts)
}

func CreateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var contact Contact
	_ = json.NewDecoder(r.Body).Decode(&contact)
	contact.ID = "1a5071bd-2960-4829-8adc-593e216b2de5"
	contact.CreatedAt = time.Now()
	contact.UpdatedAt = time.Now()
	contacts = append(contacts, contact)
	json.NewEncoder(w).Encode(contact)
}

func UpdateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, contact := range contacts {
		if contact.ID == params["id"] {
			contacts = append(contacts[:index], contacts[index+1:]...)
			var updatedContact Contact
			_ = json.NewDecoder(r.Body).Decode(&updatedContact)
			updatedContact.ID = params["id"]
			updatedContact.CreatedAt = contact.CreatedAt
			updatedContact.UpdatedAt = time.Now()
			contacts = append(contacts, updatedContact)
			json.NewEncoder(w).Encode(updatedContact)
			return
		}
	}
	json.NewEncoder(w).Encode(nil)
}

func DeleteContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, contact := range contacts {
		if contact.ID == params["id"] {
			contacts = append(contacts[:index], contacts[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(contacts)
}

func main() {
	router := mux.NewRouter()

	contacts = append(contacts, Contact{
		ID:        "1a5071bd-2960-4829-8adc-593e216b2de5",
		Name:      "fulan",
		Gender:    "male",
		Phone:     "628123456789",
		Email:     "fulan@example.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	router.HandleFunc("/contacts", GetContacts).Methods("GET")
	router.HandleFunc("/contacts", CreateContact).Methods("POST")
	router.HandleFunc("/contacts/{id}", UpdateContact).Methods("PUT")
	router.HandleFunc("/contacts/{id}", DeleteContact).Methods("DELETE")

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
