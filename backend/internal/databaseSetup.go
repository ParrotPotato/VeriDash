package internal

import (
    "github.com/dashboard/backend/database"
    "github.com/dashboard/backend/models"
)


func SetupDatabase(){
    // connect to the main database
    main := database.GetConnection("test/main.db")
    
    // check if there are tables in the main database

    main.Connection.AutoMigrate(&models.Project{})
    main.Connection.AutoMigrate(&models.TestSuite{})
    main.Connection.AutoMigrate(&models.TestRun{})
    main.Connection.AutoMigrate(&models.TestCase{})
    main.Connection.AutoMigrate(&models.TestResult{})

    // create the database corresponding to the projects 
    // which are present in the database
    // iterate over each of the projects and create the 
    // required tables 
}

