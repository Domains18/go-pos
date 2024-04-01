-- ./migrations/XXXXXXXXXX_init_schema.up.sql

-- Create users table
CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE
    );

-- Create products table
CREATE TABLE IF NOT EXISTS products (
                                        id SERIAL PRIMARY KEY,
                                        sku VARCHAR(50) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    tax_rate DECIMAL(5, 2) NOT NULL,
    taxed_price DECIMAL(10, 2) NOT NULL,
    image TEXT,
    product_category VARCHAR(100),
    sale_price DECIMAL(10, 2),
    taxed_sale_price DECIMAL(10, 2),
    quantity INT,
    available BOOLEAN DEFAULT true
    );

-- Create customers table
CREATE TABLE IF NOT EXISTS customers (
                                         id SERIAL PRIMARY KEY,
                                         name VARCHAR(255) NOT NULL,
    mobile_no VARCHAR(20),
    email VARCHAR(255) NOT NULL UNIQUE
    );

-- Create clients table for invoices
CREATE TABLE IF NOT EXISTS clients (
                                       id SERIAL PRIMARY KEY,
                                       name VARCHAR(255) NOT NULL,
    mobile_no VARCHAR(20),
    email VARCHAR(255) NOT NULL UNIQUE
    );

-- Create sales_invoice_items table
CREATE TABLE IF NOT EXISTS sales_invoice_items (
                                                   id SERIAL PRIMARY KEY,
                                                   invoice_id INT NOT NULL,
                                                   product_sku VARCHAR(50) NOT NULL,
    quantity INT NOT NULL,
    selling_price DECIMAL(10, 2) NOT NULL,
    taxed_selling_price DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (invoice_id) REFERENCES invoices(id)
    );

-- Create invoices table
CREATE TABLE IF NOT EXISTS invoices (
                                        id SERIAL PRIMARY KEY,
                                        client_id INT NOT NULL,
                                        total_amount DECIMAL(10, 2) NOT NULL,
    invoice_date DATE NOT NULL,
    status VARCHAR(20) DEFAULT 'pending',
    FOREIGN KEY (client_id) REFERENCES clients(id)
    );
