<img width="1536" height="1024" alt="SCD Golang Application Architecture Overview" src="https://github.com/user-attachments/assets/8c903f24-34d5-4a80-b50b-bed98b6b216c" />
# ⚙️ SCD Work Trial – Golang Implementation

This repository contains a solution for the **SCD (Slowly Changing Data)** Work Trial using **Golang**, **GORM**, and **SQLite**. The project demonstrates versioned data modeling, business logic separation, and unit test coverage using seeded sample data.

---

## 📌 Table of Contents

- [🧠 Problem Statement](#-problem-statement)
- [📂 Project Structure](#-project-structure)
- [🛠️ Tech Stack](#️-tech-stack)
- [🧱 Architecture Overview](#-architecture-overview)
- [⚙️ Setup Instructions](#️-setup-instructions)
- [🧪 How to Run & Test](#-how-to-run--test)
- [✅ Features Implemented](#-features-implemented)
- [📸 Loom Video Walkthrough](#-loom-video-walkthrough)

---

## 🧠 Problem Statement

Design and implement a backend system using **versioned data** where entities like `Job`, `Timelog`, and `PaymentLineItem` have versioning capabilities. Build clean abstractions that fetch the **latest versions** of the data and support testing.

---

## 📂 Project Structure

```bash
scd-golang-trial/
│
├── main.go                         # Entry point: initializes DB and prints results
├── go.mod                          # Module definition
├── go.sum                          # Dependencies checksum
│
├── models/
│   ├── models.go                   # GORM model definitions (Job, Timelog, PaymentLineItem)
│
├── scd/
│   ├── abstraction.go              # Business logic to get latest versions
│   └── abstraction_test.go        # Unit tests using sample data
│
└── scd.db                          # SQLite database (auto-generated)



🛠️ Tech Stack
Language: Golang v1.21+

ORM: GORM

Database: SQLite

Testing: Go's built-in testing package

Diagram Tool: Eraser.io

🧱 Architecture Overview


Key Components:
Main Function:

Initializes SQLite DB using GORM

Seeds sample data with multiple versions

Prints output to console

Models Package:

Defines 3 versioned entities: Job, Timelog, PaymentLineItem

Uses composite primary key (id + version)

Business Logic:

scd.GetLatestJobsForCompany returns latest job version

scd.GetLatestTimelogsForJob

scd.GetLatestPaymentItemsForJob

Persistence Layer:

GORM handles DB operations

SQLite used for local testing and simplicity

Testing:

Unit tests verify that only latest versions are returned

⚙️ Setup Instructions
🖥️ Prerequisites
Golang v1.21+

GCC (required for go-sqlite3)

Git

🧰 Clone the Repository

git clone https://github.com/DevOpsatishmahakalkar/scd-golang-trial.git
cd scd-golang-trial
📦 Install Dependencies

go mod tidy
🚀 Run the Application

go run main.go
Expected output:

✅ Latest Jobs for Company 'comp1':
{ID:job1 Version:2 UID:job_uid_v2 Status:active Rate:22 Title:SDE CompanyID:comp1 ContractorID:cont1}
🔹 Latest Timelogs for Job: job_uid_v2
{ID:tl1 Version:1 UID:tl_uid_1 Duration:1200 ...}
💰 Latest Payment Line Items for Job: job_uid_v2

How to Run Unit Tests

go test ./scd

Expected output:

PASS
ok      scd-golang-trial/scd  0.123s


✅ Features Implemented
 GORM model definitions with versioned IDs

 SQLite DB with seed data

 Business logic abstractions to return latest records

 Console output for result visibility

 Unit tests for business logic validation
