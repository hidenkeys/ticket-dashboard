
# Ticket Monitoring Dashboard - README

## Overview

This project provides a **generic backend** for building **ticket/issue monitoring dashboards**. The system allows you to track and monitor the progress of tickets/issues through various stages, from creation to resolution. It is designed to be flexible and scalable, enabling users to extend it to fit their own needs.

Whether you are building a **helpdesk system**, **workflow management dashboard**, or simply need a way to track the status of issues, this setup provides the necessary infrastructure for creating a fully-fledged dashboard.

## Features

* **Ticket Management**: Track tickets/issues through different stages.
* **Stage Tracking**: Each ticket can move through multiple stages, each having its own start time, end time, type, and description.
* **Workflow Monitoring**: Easily monitor where each issue is in the workflow, and identify stages where issues are pending or delayed.
* **Generic Structure**: This system is not tied to any specific use case, making it versatile and customizable for different needs.
* **Self-Referencing Stages**: Major stages can have related minor stages, enabling detailed tracking of complex workflows.
* **Extensible**: Add new stages, customize ticket fields, and build your custom views on top of this base model.

## Models

There are two primary models in this system:

### 1. **Ticket Model**

The `Ticket` model represents an issue or a task that needs to be tracked through various stages.

Fields:

* `id`: Primary key
* `customer_name`: Name of the customer or organization associated with the ticket
* `helpdeskurl`: URL to the helpdesk or related support page
* `current_stage_id`: The ID of the current stage the ticket is in
* `stages`: A many-to-many relationship with the `Stage` model (tickets can have multiple stages)
* `status`: Current status of the ticket (e.g., open, closed, pending, etc.)
* `contact_person`: Person handling the ticket or issue
* `project`: The project associated with the ticket

### 2. **Stage Model**

The `Stage` model represents each stage a ticket moves through. Stages can be either major or minor, with minor stages being related to a major stage.

Fields:

* `id`: Primary key
* `stage_start_time`: The start time of the stage
* `stage_end_time`: The end time of the stage
* `stage_type`: Whether the stage is a major or minor stage
* `major_stage_id`: The ID of the major stage (if the current stage is a minor stage)
* `stage_description`: An optional description of the stage
* `contact_emails`: A list of emails associated with the stage (for notifications or updates)

### Relationships:

* **Ticket -> Stage**: A `Ticket` has a foreign key `current_stage_id` pointing to the `Stage` model.
* **Stage -> Stage**: A `Stage` can have a reference to another `Stage` if it is a minor stage (via `major_stage_id`), allowing for detailed workflow tracking.

## API Endpoints

Below are the general endpoints for interacting with this system.

### **1. GET /tickets**

Fetches all tickets in the system.

* **Response**: List of all tickets.

### **2. GET /tickets/\:id**

Fetches a single ticket by its ID.

* **Response**: Ticket details, including its current stage and all stages it has gone through.

### **3. POST /tickets**

Creates a new ticket in the system.

* **Request Body**: JSON object with ticket details (e.g., customer name, project, status, etc.)

### **4. PUT /tickets/\:id**

Updates an existing ticket’s information.

* **Request Body**: JSON object with fields to update (e.g., status, stage, etc.)

### **5. GET /stages**

Fetches all stages in the system.

* **Response**: List of all stages, including their type (major/minor), start time, and end time.

### **6. GET /stages/\:id**

Fetches a single stage by its ID.

* **Response**: Stage details, including its description and the related major stage if applicable.

### **7. POST /stages**

Creates a new stage in the system.

* **Request Body**: JSON object with stage details (e.g., type, start time, end time, description, etc.)

### **8. PUT /stages/\:id**

Updates an existing stage.

* **Request Body**: JSON object with fields to update (e.g., stage type, start time, end time, etc.)

## Setup Instructions

### Requirements:

* Go 1.16+
* Database (MySQL/PostgreSQL recommended)

### Steps to Set Up:

1. **Clone the repository**:

   ```bash
   git clone https://github.com/hidenkeys/ticket-monitoring-dashboard.git
   cd ticket-monitoring-dashboard
   ```

2. **Install dependencies**:
   Run the following command to install required Go packages:

   ```bash
   go mod tidy
   ```

3. **Set up your database**:

   * Create a new database in your DBMS (e.g., MySQL or PostgreSQL).
   * Update your database connection settings in the `main.go` file.

4. **Run the application**:

   ```bash
   go run main.go
   ```

5. **Access the API**:
   Once the application is running, you can start making requests to the endpoints defined above.

## Extending the System

This system is built to be flexible, so you can easily extend it to suit your needs. Here are a few ideas for extensions:

* **Custom Views**: Build custom frontend views to visualize the progress of tickets and issues.
* **Email Notifications**: Integrate with an email service to send notifications to users when a ticket is updated or reaches a certain stage.
* **Permission Management**: Implement role-based access control (RBAC) for different users (e.g., admin, support staff, customers).
* **Advanced Search**: Add search functionality to allow users to search tickets by status, customer, project, or stage.

