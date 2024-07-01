package main

import (
    "github.com/dashboard/backend/internal"

)


func main() {
    internal.SetupDatabase()
    internal.StartServer()
}
