# CRM Backend System

## Overview

This CRM (Customer Relationship Management) backend system is designed to handle customer interactions, manage email notifications, track email opens, and manage user notifications for tasks, meetings, and follow-ups. The system is built using Go and Gin framework, with MongoDB as the database.

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

