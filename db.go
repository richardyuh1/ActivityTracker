package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DBClient struct {
	db *sql.DB
}

func newDBClient(
	dbName string,
	userName string,
	password string,
) (*DBClient, error) {
	dataSourceName := fmt.Sprintf("%s:%s@/%s", userName, password, dbName)
	fmt.Printf("Data source name: %v\n", dataSourceName)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("SQL open error found: %v\n", err.Error())
	}
	return &DBClient{db: db}, nil
}

func (d *DBClient) Ping() error {
	return d.db.Ping()
}

func (d *DBClient) Close() error {
	return d.db.Close()
}

func (d *DBClient) CreateActivitiesTable() error {
	// Make id a UUID
	_, err := d.db.Exec("CREATE TABLE IF NOT EXISTS Activities " +
		"(id INT PRIMARY KEY AUTO_INCREMENT, activity_name VARCHAR(32), " +
		"activity_time DATETIME, duration INT, notes VARCHAR(256))")
	return err
}

func (d *DBClient) AddActivity(
	activityName string,
	activityTime string,
	duration int,
	notes string,
) error {
	_, err := d.db.Exec(
		"INSERT INTO Activities "+
			"(activity_name, activity_time, duration, notes) "+
			"VALUES(?, ?, ?, ?);",
		activityName,
		activityTime,
		duration,
		notes,
	)
	return err
}

func (d *DBClient) GetActivities() (*sql.Rows, error) {
	rows, err := d.db.Query("select * from Activities;")
	return rows, err
}

func (d *DBClient) UpdateActivity(
	id string,
	activityName string,
	activityTime string,
	duration int,
	notes string,
) error {
	res, err := d.db.Exec(
		"UPDATE Activities SET activity_name = ?, activity_time = ?, "+
			"duration = ?, notes = ? WHERE id = ?",
		activityName,
		activityTime,
		duration,
		notes,
		id,
	)
	rowsAff, _ := res.RowsAffected()
	if err != nil || rowsAff != 1 {
		return fmt.Errorf(
			"Error updating activity: err: %v, rows affected: %v",
			err,
			rowsAff,
		)
	}
	return nil
}

func (d *DBClient) DeleteActivity(id string) error {
	res, err := d.db.Exec("DELETE FROM Activities WHERE id = ?;", id)
	rowsAff, _ := res.RowsAffected()
	if err != nil || rowsAff != 1 {
		return fmt.Errorf(
			"Error deleting activity: err: %v, rows affected: %v",
			err,
			rowsAff,
		)
	}
	return nil
}
