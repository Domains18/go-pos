
## SaaS Project using Golang and WhatsApp Cloud API

This project is a SaaS platform that acts as a middleware between businesses and the WhatsApp Cloud API. It allows businesses to:

- Register with 3 admin users
- Add payment methods
- Create a profile
- Manage WhatsApp numbers
- Utilize the WhatsApp Cloud API functionalities (future implementation)

This README serves as a guide for developers contributing to the project.

### Prerequisites

* Golang (version 1.18 or higher recommended)
* Basic understanding of Golang concepts (concurrency, channels, etc.)
* Familiarity with relational databases (e.g., PostgreSQL)
* Experience with cloud platforms (optional, but beneficial)

### Getting Started

1. **Clone the repository:**

```bash {"id":"01HSANE19M69WVJ4T9QVG2BEDW"}
git clone https://github.com/Nerds-Catapult/notiwa.git
```

2. **Install dependencies:**

```bash {"id":"01HSANE19M69WVJ4T9QY2MYK3M"}
go mod download
```

3. **Configure environment variables:**

Create a `.env` file in the project root directory with the following variables:

- `DATABASE_URL`: Connection string for your database
- `SECRET_KEY`: Secret key for JWT token generation

4. **Run the application:**

```bash {"id":"01HSANE19M69WVJ4T9QZ2VB76E"}
go run main.go
```

### Project Structure

The project is organized into the following directories:

- **cmd:** Contains the main application entry point (`main.go`).
- **config:** Stores configuration files (e.g., database connection details).
- **internal:** Houses core application logic:
   - **models:** Defines data structures for users, businesses, etc.
   - **services:** Encapsulates business logic for user registration, payment processing, profile management, etc.
   - **repositories:** Provides interfaces and implementations for interacting with the database.
   - **utils:** Contains helper functions and utilities used throughout the project.

- **tests:** Includes unit and integration tests for various components.

### Contributing

We welcome contributions from the developer community! Here's how you can contribute:

1. **Fork the repository:** Create a fork of this repository on your GitHub account.
2. **Create a branch:** Create a new branch for your feature or bug fix.
3. **Develop your changes:** Make changes to the codebase following the existing coding style and conventions.
4. **Write unit tests:** Ensure your changes are covered by unit tests.
5. **Commit your changes:** Commit your changes with a clear and concise message.
6. **Push your branch:** Push your branch to your forked repository.
7. **Create a pull request:** Create a pull request to merge your changes into the main repository.

### Code Style

We follow a consistent code style for better readability and maintainability. Use tools like `go fmt` and `go fmt` to ensure proper formatting.

### Documentation

API documentation will be added progressively as features are developed.

### Testing

Unit and integration tests are crucial for ensuring code quality. We encourage writing tests for your contributions.

### Deployment

Deployment instructions and environment configurations will be detailed as the project matures.

### License

This project is licensed under the MIT License. See the `LICENSE` file for details.

### Additional Resources

- Golang Documentation: [https://go.dev/doc/](https://go.dev/doc/)
- WhatsApp Cloud API: [https://developers.facebook.com/products/whatsapp/](https://developers.facebook.com/products/whatsapp/)

