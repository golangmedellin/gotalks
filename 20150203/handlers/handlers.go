package handlers

import (
    "encoding/json"
    "github.com/sescobb27/meetup/models"
    "net/http"
)

func GetCategories(res http.ResponseWriter, req *http.Request) {
    res.Header().Set("Content-Type", "application/json")
    categories := models.GetCategories()
    json_categories, err := json.Marshal(categories)

    if err != nil {
        panic(err)
    }

    res.Write(json_categories)
}

func GetLocations(res http.ResponseWriter, req *http.Request) {
    res.Header().Set("Content-Type", "application/json")
    locations := models.GetLocations()
    json_locations, err := json.Marshal(locations)

    if err != nil {
        panic(err)
    }

    res.Write(json_locations)
}
