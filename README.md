# CRM Backend System

## Overview

This CRM (Customer Relationship Management) backend system is designed to handle customer interactions, manage email notifications, track email opens, and manage user notifications for tasks, meetings, and follow-ups. The system is built using Go and Gin framework, with MongoDB as the database.

---

## Hosted Link
- [Live](http://67.205.140.148:8080/)
- [API Documentation hosted link](https://documenter.getpostman.com/view/14938340/2sA3Qs9BmN)

---

## API Design

```
 POST   /register               
 POST   /login                   
 GET    /admin/users             
 GET    /admin/users/:id         
 PUT    /admin/users/:id         
 DELETE /admin/users/:id         
 POST   /admin/tickets            
 POST   /admin/tickets            
 PUT    /admin/tickets/:id/resolve 
 POST   /admin/interactions       
 GET    /admin/interactions/customer/:customer_id 
 GET    /admin/reports/customer-interactions 
 POST   /admin/send-email         
 GET    /admin/track/open/:trackingID 
 POST   /admin/notifications      
 GET    /admin/notifications      
```

---

## Database Schema Design

### Collections

### 1. Users
- **users**
```json
{
  "_id": "ObjectId",
  "name": "string",
  "email": "string",
  "password": "string",
  "company": "string",
  "status": "string",
  "notes": "string"
}
```

### 2. Notifications

```json
{
  "_id": "ObjectId",
  "user_id": "ObjectId",
  "message": "string",
  "created_at": "datetime",
  "notified_at": "datetime",
  "is_notified": "boolean",
  "event_time": "datetime"
}

```

### 3. Emails

```json
{
  "_id": "ObjectId",
  "from": "string",
  "to": "string",
  "subject": "string",
  "body": "string",
  "created_at": "datetime",
  "tracking_id": "string"
}
```


### 4. Email Opens

```json
{
  "_id": "ObjectId",
  "from": "string",
  "to": "string",
  "subject": "string",
  "status": "string"
}
"status": "string" // e.g., opened, sent
}
```


### 5. Interactions
```json
{
  "_id": "ObjectId",
  "customer_id": "ObjectId",
  "user_id": "ObjectId",
  "title": "string",
  "details": "string",
  "scheduled_at": "datetime",
  "created_at": "datetime"
}
```

### 6. Tickets

```json
{
  "_id": "ObjectId",
  "customer_id": "ObjectId",
  "title": "string",
  "description": "string",
  "status": "string",
  "created_at": "datetime",
  "resolved_at": "datetime"
}
```

---

## Code Structure

```bash
.
├── README.md
├── config
│   └── config.go
├── controllers
│   ├── emailController.go
│   ├── interactionController.go
│   ├── notification.go
│   ├── reportController.go
│   ├── ticketController.go
│   └── userController.go
├── go.mod
├── go.sum
├── main.go
├── middleware
│   └── authMiddleware.go
├── models
│   ├── emailTracking.go
│   ├── event.go
│   ├── interactionModel.go
│   ├── notification.go
│   ├── report.go
│   ├── ticketModel.go
│   ├── user.go
│   └── userModel.go
├── routes
│   └── routes.go
├── services
│   └── useService.go
└── utils
    ├── email.go
    ├── env.go
    ├── hash.go
    ├── jwt.go
    └── notification.go
```

- `controllers`: Contains the controller files for handling the business logic.
- `routes`: Contains the routes for the application.
- `models`: Contains the models for the application.
- `services`: Contains the services for the application.
- `utils`: Contains the utility functions for the application.
- `middleware`: Contains the middleware functions for the application.
- `config`: Contains the configuration for the db
- `main.go`: Entry point of the application.

---

### How to run

```bash
git clone https://github.com/d-cryptic/crm-golang-backend.git
cd crm-golang-backend
go mod tidy
export MONGO_URI=
export JWT_SECRET=
export SMTP_HOST=smtp.gmail.com
export SMTP_PORT=465
export SMTP_USERNAME=example@gmail.com
export SMTP_PASSWORD=
export PORT=8000
export SERVER_URL=http://localhost:8000
go run main.go
go tests ./...
docker build -t crm-backend .
docker run -p 8080:8080 crm-backend
```