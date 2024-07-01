package internal 

import (
    "github.com/dashboard/backend/api"
    "github.com/gin-gonic/gin"
)

func StartServer(){
    router := gin.Default()
    router.GET("/api/project_name_id/", api.GetProjectNameAndID)
    router.Run(":8080")
}
