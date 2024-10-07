# ONLINE STORE APP API

This is the Online Store REST API documentation, this API is used to log in and register users, besides that it is used to retrieve products by category, add products to the cart, view the cart, delete product in cart, and make transactions.


## Table of Contents
- [Documentation](#documentation)
- [Tech Stack](#tech-stack)
- [Features](#features)
- [Installation](#installation)

## Documentation

### Entity Relationship Diagram (ERD)
![image](https://github.com/MuhammadIbraAlfathar/assets/blob/main/Untitled%20(1).png?raw=true)



## Tech Stack
- **Golang**: The programming language used for the core backend logic.
- **Gin**: A high-performance HTTP web framework for building the RESTful API.
- **GORM**: An ORM library for Golang, used to interact with the PostgreSQL database.
- **PostgreSQL**: The primary database for storing all application data.
- **Docker**: Containerization for easy deployment and management.
- **JWT**: For handling user authentication and authorization.
- **Github Actions**: For build and deploy image docker.

## Features

- **User Authentication**: Secure authentication using JWT.
- **Cart**: Add product to cart, view cart, and delete product in cart.
- **Transaction**: Make transaction.
- **Product**: View product by category.

## Installation

Before you begin, ensure you have the these installed on your machine:
- Docker

### Steps
1. **Clone the Repository:**
    ```bash
    https://github.com/MuhammadIbraAlfathar/online-store-app.git
   ```

2. **Navigate to the Project Directory:**
    ```bash
    cd online-store-app
    ```

3. **Set Up Environment Variables:**  
   Create a `.env` file in the root directory and provide the necessary environment variables. See `.env.example` file for reference.

4. **Start the Server:**
   ```bash
   docker compose up
   ```