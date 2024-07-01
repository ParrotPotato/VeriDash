package models

import (
    "gorm.io/gorm"
)

type Project struct {
    gorm.Model
	Name        string
	Description string


    // gorm has many association
    TestSuite  []TestSuite
    TestCase   []TestCase
    TestRun    []TestRun
    TestResult []TestResult
}
