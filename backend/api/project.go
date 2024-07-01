package api

import(
    "github.com/dashboard/backend/database"
    "github.com/dashboard/backend/models"
    "github.com/gin-gonic/gin"

    "net/http"
    "fmt"
)

func GetProjectNameAndID(context * gin.Context){
    db := database.GetConnection("test/main.db")
    var projects []models.Project
    db.Connection.Find(&projects)
    project_id_name := make([]map[string]interface{}, 0)
    for _, project := range projects {
        entry :=  map[string]interface{}{"Id": fmt.Sprintf("%d", project.ID), "Name": project.Name}
        project_id_name = append(project_id_name, entry)
    }
    context.IndentedJSON(http.StatusOK, project_id_name)
}
