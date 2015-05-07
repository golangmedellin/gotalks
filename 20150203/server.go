package main

import (
    "flag"
    . "github.com/sescobb27/meetup/handlers"
    "github.com/sescobb27/meetup/models"
    "net/http"
    "os"
)

var (
    categories = []string{
        "Pasta",
        "Carne",
        "Pollo",
        "Ensalada",
        "Desayuno",
        "Almuerzo",
        "Postre",
        "Sopa",
        "Vegetariana",
        "Menu Infantil",
        "Comida Rapida",
        "Almuerzo para 2",
        "Desayuno para 2",
        "Comida para 2",
        "Ensalada de Frutas",
        "Gourmet",
    }
    descriptions = []string{
        "Ricas Pasta",
        "Ricas Carnes",
        "Rico Pollo",
        "Ricas Ensaladas",
        "Ricos Desayunos",
        "Ricos Almuerzos",
        "Ricos Postres",
        "Ricas Sopas",
        "Rica Comida Vegetariana",
        "Ricos Menu Infantil",
        "Rica Comida Rapida",
        "Ricos Almuerzo para 2",
        "Ricos Desayuno para 2",
        "Ricas Comidas para 2",
        "Ricas Ensaladas de Frutas",
        "Ricas Comida Gourmet",
    }
    locations = []string{
        "Poblado",
        "Laureles",
        "Envigado",
        "Caldas",
        "Sabaneta",
        "Colores",
        "Estadio",
        "Calazans",
        "Bello",
        "Boston",
        "Prado Centro",
        "Itagui",
        "Belen",
        "Guayabal",
    }
)

func insert_Locations() {
    for _, location := range locations {
        l := &models.Location{Name: location}
        l.Create()
    }
}
func insert_Categories() {
    for i, category := range categories {
        c := &models.Category{
            Name:        category,
            Description: descriptions[i],
        }
        c.Create()
    }
}

func Seed() {
    insert_Categories()
    insert_Locations()
}

func main() {
    seed := flag.Bool("seed", false, "Seed the database")
    flag.Parse()

    if *seed {
        Seed()
        os.Exit(0)
    }
    server := http.NewServeMux()
    server.Handle("/categories", Get(GetCategories))
    server.Handle("/locations", Get(GetLocations))
    http.ListenAndServe(":3000", server)
}
