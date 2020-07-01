package main

import (
	_ "github.com/guypeled76/go-bigquery-driver/gorm/dialect"
	"github.com/jinzhu/gorm"
	"log"
)

type RunTestProject struct {
	Name string
}

func main() {
	db, err := gorm.Open("bigquery", "bigquery://unity-rd-perf-test-data-prd/location/perf_test_results")
	if err != nil {
		log.Fatal(err)
	}

	var projects []RunTestProject

	err = db.Not("Name", "").Find(RunTestProject{}).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, project := range projects {
		log.Println(project)
	}

	defer db.Close()
	// Do Something with the DB

}