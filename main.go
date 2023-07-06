package main

import (
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Start main execution")
	fileName := "../test.csv"
	csvReader := newCSVReader(fileName)
	contents, err := csvReader.readAll()
	if err != nil {
		fmt.Printf("Error found: %v\n", err.Error())
		return
	}
	fmt.Printf("Contents: %v\n", contents)
	headers, err := csvReader.Headers()
	if err != nil {
		fmt.Printf("Error found: %v\n", err.Error())
		return
	}
	fmt.Printf("Headers: %v\n", headers)
	rows, err := csvReader.Rows()
	if err != nil {
		fmt.Printf("Error found: %v\n", err.Error())
		return
	}
	fmt.Printf("Rows: %v\n", rows)
	cols, err := csvReader.Columns()
	if err != nil {
		fmt.Printf("Error found: %v\n", err.Error())
		return
	}
	fmt.Printf("Cols: %v\n", cols)
	userName := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbClient, err := newDBClient(dbName, userName, password)
	if err != nil {
		fmt.Printf("SQL Client error: %v\n", err.Error())
		return
	}
	defer dbClient.Close()

	err = dbClient.Ping()
	if err != nil {
		fmt.Printf("Ping error found: %v\n", err.Error())
		return
	}
	fmt.Println("Successful db connection!")
	err = dbClient.CreateActivitiesTable()
	if err != nil {
		fmt.Printf("Create table error found: %v\n", err.Error())
		return
	}
	activityName := "Bike"
	activityTime := time.Now().UTC().Format("2006-01-02T15:04:05")
	duration := 3600
	notes := "SF Loop"
	err = dbClient.AddActivity(activityName, activityTime, duration, notes)
	if err != nil {
		fmt.Printf("Add activity error found: %v\n", err.Error())
	} else {
		fmt.Println("Successful add activity operation")
	}
	queryRows, err := dbClient.GetActivities()
	defer queryRows.Close()
	if err != nil {
		fmt.Printf("Get activities error found: %v\n", err.Error())
	} else {
		fmt.Printf("Activities rows: %v\n", queryRows)
	}

	err = dbClient.UpdateActivity("1", "bike", activityTime, 1000, "training")
	if err != nil {
		fmt.Printf("Update activity error found: %v\n", err.Error())
	} else {
		fmt.Println("Successful update operation")
	}

	err = dbClient.DeleteActivity("14")
	if err != nil {
		fmt.Printf("Delete activity error found: %v\n", err.Error())
	} else {
		fmt.Println("Successful delete operation")
	}
}
