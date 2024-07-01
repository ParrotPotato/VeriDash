package models 

import (
    "gorm.io/gorm"
)


// TestCase is one of the test scenarios to be tested
type TestCase struct {
    gorm.Model

    ProjectID   uint
    TestSuiteID uint

    // test is the history of execution of this test cases
    TestResult  []TestResult
}
