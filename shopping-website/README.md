# Shopping Website

This is a shopping website project built using Go and the Gin framework. The application provides functionalities for both registered and unregistered users, including user authentication, product management, shopping cart operations, and order processing.

## Features

### Unregistered Users
1. **Member Registration**: Users can create a new account.
2. **Member Login**: Users can log in to their accounts.
3. **Product Query**: Users can browse available products.
4. **View Product Details**: Users can view detailed information about a specific product.
5. **Add to Cart**: Users can add products to their shopping cart.

### Registered Users
1. **Member Modification**: Users can update their account information.
2. **Member Logout**: Users can log out of their accounts.
3. **View/Modify Cart**: Users can view and modify their shopping cart.
4. **Checkout**: Users can complete their purchase.
5. **View Order History**: Users can view their past orders.
6. **View Order Details**: Users can view details of a specific order.
7. **Notify Transfer**: Users can notify the system of a completed transfer.

## Project Structure

```
shopping-website
├── controllers
│   ├── auth.go
│   ├── product.go
│   ├── cart.go
│   ├── order.go
├── models
│   ├── user.go
│   ├── product.go
│   ├── cart.go
│   ├── order.go
├── routes
│   ├── auth.go
│   ├── product.go
│   ├── cart.go
│   ├── order.go
├── services
│   ├── auth_service.go
│   ├── product_service.go
│   ├── cart_service.go
│   ├── order_service.go
├── utils
│   └── utils.go
├── main.go
├── go.mod
└── README.md
```

## Setup Instructions

1. Clone the repository:
   ```
   git clone <repository-url>
   cd shopping-website
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Run the application:
   ```
   go run main.go
   ```

4. Access the application at `http://localhost:8080`.

## Usage Guidelines

- Follow the API documentation for endpoint usage.
- Ensure to handle user authentication properly.
- Validate user inputs to maintain data integrity.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.