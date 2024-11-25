# RESTful Inventory Management System with Golang

This project is a RESTful inventory management system developed using Go Fiber and MongoDB. It provides efficient Create, Read, Update, and Delete (CRUD) operations to manage product inventories with ease.

The backend leverages the lightweight and high-performance Go Fiber framework, ensuring quick response times and efficient resource handling. MongoDB serves as the database, offering a flexible and scalable solution for storing product data. Additionally, the system features a Text User Interface (TUI) to fetch and display all product details directly in a console-based environment, enhancing accessibility for users who prefer terminal-based interactions.

The system is designed to be modular, maintainable, and scalable, making it an ideal choice for inventory management needs in diverse environments.

## Installation

Clone the repository

```bash
git clone https://github.com/shohan-pherones/inventory-management-system.git
```

Navigate into the project directory

```bash
cd inventory-management-system
```

Install the required packages:

```bash
go mod tidy
```

Create a `.env` file in the project root directory and define the following variables:

```bash
MONGO_URI
DATABASE_NAME
```

Launch the application using:

```bash
go run main.go
```

Your application should now be running locally at `http://127.0.0.1:4000`
