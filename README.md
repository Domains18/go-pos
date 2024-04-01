# Notiwa: SaaS Platform for WhatsApp Integration

Notiwa is a Software-as-a-Service (SaaS) platform that acts as a middleware between businesses and the WhatsApp Cloud API. It enables businesses to seamlessly integrate WhatsApp messaging into their operations.

## Features

1. **User Registration and Authentication**:
  - Businesses can register with Notiwa using three admin users.
  - Secure authentication using JWT tokens.

2. **Payment Methods**:
  - Businesses can add and manage payment methods for subscription billing.

3. **Business Profile Management**:
  - Create and update business profiles.
  - Store essential information such as company name, contact details, and business description.

4. **WhatsApp Number Management**:
  - Businesses can associate WhatsApp numbers with their accounts.
  - Manage multiple WhatsApp numbers for different purposes (e.g., customer support, notifications).

5. **WhatsApp Cloud API Integration (Future Implementation)**:
  - Utilize the WhatsApp Cloud API for sending messages, notifications, and automating workflows.

## Prerequisites

Before getting started, ensure you have the following:

- Golang (version 1.18 or higher recommended)
- Basic understanding of Golang concepts (concurrency, channels, etc.)
- Familiarity with relational databases (e.g., PostgreSQL)
- Optional: Experience with cloud platforms (for deployment)

## Getting Started

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/Nerds-Catapult/notiwa.git
   ```

2. **Install Dependencies:**
   ```bash
   go mod download
   ```

3. **Configure Environment Variables:**

   Create a `.env` file in the project root directory with the following variables:
  - `DATABASE_URL`: Connection string for your PostgreSQL database
  - `SECRET_KEY`: Secret key for JWT token generation

4. **Run the Application:**
   ```bash
   go run main.go
   ```

## Project Structure

The Notiwa project follows a well-organized structure:

- **`cmd`**: Contains the main application entry point (`main.go`).
- **`config`**: Stores configuration files (e.g., database connection details).
- **`internal`**: Core application logic resides here:
  - **`models`**: Defines data structures (e.g., users, business profiles).
  - **`services`**: Encapsulates business logic (e.g., user registration, payment processing).
  - **`repositories`**: Provides database interaction implementations.
  - **`utils`**: Contains helper functions and utilities.

## Contributing

Contributions to Notiwa are welcome! Please follow the guidelines in the [CONTRIBUTING.md](CONTRIBUTING.md) file.

## License

Notiwa is open-source software released under the [MIT License](LICENSE).

