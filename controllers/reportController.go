package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/d-cryptic/crm-golang-backend/models"
)

func GenerateCustomerInteractionsReport(c *gin.Context) {
    // Fetch data from the database or any other source
    interactionsData := fetchCustomerInteractionsData()

    // Prepare data for Chart.js
    var labels []string
    var values []int
    for _, data := range interactionsData {
        labels = append(labels, data.Month)
        values = append(values, data.Interactions)
    }

    // Render the chart using Chart.js format
    chartData := gin.H{
        "labels": labels,
        "datasets": []gin.H{
            {
                "label": "Customer Interactions",
                "data":  values,
                "backgroundColor": "rgba(75, 192, 192, 0.2)",
                "borderColor": "rgba(75, 192, 192, 1)",
                "borderWidth": 1,
            },
        },
    }

    // Return the chart data as JSON
    c.JSON(http.StatusOK, chartData)
}

// Mock function to fetch customer interactions data (replace with actual database query)
func fetchCustomerInteractionsData() []models.CustomerInteraction {
    // Here you would typically fetch data from your database
    // For demo purposes, let's mock some data
    return []models.CustomerInteraction{
        {Month: "January", Interactions: 10},
        {Month: "February", Interactions: 20},
        {Month: "March", Interactions: 15},
        {Month: "April", Interactions: 25},
        {Month: "May", Interactions: 30},
    }
}
