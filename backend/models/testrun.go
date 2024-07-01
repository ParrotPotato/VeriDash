package models
import (
    "gorm.io/gorm"
)


// TestRun belong to a Project and happenes as in a test suite in execution
type TestRun struct {
    gorm.Model
    
    ProjectID   uint
    TestSuiteID uint 

    // this is the test result which happend when this test run was executed
    TestResult  []TestResult
}
