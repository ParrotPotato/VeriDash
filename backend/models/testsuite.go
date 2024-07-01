package models
import (
    "gorm.io/gorm"
)

// TestSuite belong to the project and has all the associated Testcases
type TestSuite struct {
    gorm.Model

    ProjectID   uint 

    TestCase    []TestCase
    TestRun     []TestRun
    TestResult  []TestResult

    Type        string 
}
