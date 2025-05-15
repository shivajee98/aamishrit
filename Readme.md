# Aamishrit E-commerce Platform

This project is a Go-based e-commerce platform with separate functionalities for customers and administrators.

Live Website Link : https://aamishrit.com 

## Functionality

This project provides two main services:

1.  **Admin Service:**
    * Located in `cmd/admin-server`.
    * Provides administrative access to the database and other backend functionalities.
    * Currently includes routes for managing products and categories:
        * `POST /admin/products`: Create a new product.
        * `GET /admin/products/:id`: Get a product by ID.
        * `PUT /admin/products/:id`: Update an existing product.
        * `DELETE /admin/products/:id`: Delete a product.
        * `POST /admin/category`: Create a new category.
        * `GET /admin/category/:id`: Get a category by ID.
        * `PUT /admin/category/:id`: Update an existing category.
        * `DELETE /admin/category/:id`: Delete a category.
        * `GET /admin/category`: Get all categories.
    * Future enhancements may include order management, user banning, and refund APIs.

2.  **Customer Service:**
    * Located in `cmd/customer-server`.
    * Provides APIs for customers to interact with the e-commerce platform.
    * Includes the following routes under the `/api` prefix:
        * **Public Routes:**
            * `POST /register`: Register a new user.
            * `POST /login`: User login.
            * `GET /products/:id`: Get a product by ID.
            * `GET /products`: List all products.
        * **Protected Routes (requires authentication via Clerk middleware):**
            * **User:**
                * `PUT /user`: Update user information.
                * `GET /user/check`: Check user by Clerk ID.
                * `POST /user/register`: Register a new user (protected endpoint).
            * **Product:**
                * `POST /products`: Create a new product (likely for sellers/admins).
                * `PUT /products/:id`: Update a product.
                * `DELETE /products/:id`: Delete a product.
            * **Cart:**
                * `POST /cart`: Add a product to the cart.
                * `GET /cart`: Get the user's cart.
                * `DELETE /cart/:cart_id`: Remove an item from the cart.
                * `DELETE /cart/clear`: Clear the entire cart.
            * **Reviews:**
                * `POST /reviews`: Add a product review.
                * `GET /reviews/:product_id`: Get reviews for a specific product.
                * `PUT /reviews/:review_id`: Update a review.
                * `DELETE /reviews/:review_id`: Delete a review.
            * **Address:**
                * `GET /address`: Get all user addresses.
                * `POST /address`: Create a new address.
                * `GET /address/:id`: Get an address by ID.
                * `PUT /address/:id`: Update an address.
                * `DELETE /address/:id`: Delete an address.
                * `PUT /address/:id/default`: Set an address as the default.
                * `GET /address/default`: Get the default address.
        * **Orders (under `/orders`):**
            * `POST /orders`: Place a new order.
            * `GET /orders/:order_id`: Get order details.
            * `GET /orders/user/:user_id`: Get orders for a specific user.
            * `PUT /orders/:order_id`: Update order status.
            * `DELETE /orders/:order_id`: Cancel an order.
        * **Category (under `/categories`):**
            * `GET /categories`: Get all categories.
            * `GET /categories/:id`: Get a category by ID.

## Running the Project

To run this project, you need to build the binaries for both the admin and customer servers.

1.  **Build the binaries:**

    ```bash
    cd cmd/customer-server
    go build -o customer-server main.go

    cd ../admin-server
    go build -o admin-server main.go

    cd ../.. # Navigate back to the project root
    ```

2.  **Run the servers:**

    ```bash
    cd cmd/customer-server
    ./customer-server &

    cd ../admin-server
    ./admin-server &
    ```


Make sure you have Go installed and your Go environment is set up correctly.