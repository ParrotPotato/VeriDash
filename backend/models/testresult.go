package models
import (
    "gorm.io/gorm"
)


// TestResult is what happened when a test case was 
// executed
type TestResult struct{
    gorm.Model

    ProjectID   uint
    TestCaseID  uint
    TestSuiteID uint
    TestRunID   uint
}
