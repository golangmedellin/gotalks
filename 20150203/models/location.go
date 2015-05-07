package models

type Location struct {
        Id      int     `json:"id"`
        Name    string  `json:"name"`
}

func (l *Location) Create() {
        db, err := stablishConnection()
        if err != nil {
                panic(err)
        }
        defer db.Close()

        query := `INSERT INTO locations(name) VALUES ($1)`
        _, err = db.Exec(query, l.Name)

        if err != nil {
                panic(err)
        }
}

func GetLocations() []*Location {
        db, err := stablishConnection()
        if err != nil {
                panic(err)
        }
        defer db.Close()

        query := `SELECT name FROM locations`

        location_rows, err := db.Query(query)
        if err != nil {
                panic(err)
        }

        if location_rows == nil {
                panic(location_rows)
        }

        locations := []*Location{}
        for location_rows.Next() {
                location := Location{}
                err = location_rows.Scan(&location.Name)
                locations = append(locations, &location)
        }

        return locations
}
