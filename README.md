# microServices-grpc-project_product_svc
## Microservice Product-svc

This repository contains the implementation of a microservice-based product management system built using Go, gRPC, PostgreSQL, and Protocol Buffers.

### Features

- Create a new product with name, stock, and price information.
- Retrieve a product by its unique identifier (ID).
- Decrease the stock of a product when an order is placed.

### Technologies and Frameworks Used

- Go: The programming language used to develop the microservice.
- gRPC: The communication protocol used for client-server interactions.
- PostgreSQL (psql): The database used for storing product information.
- Protocol Buffers (protobuf): The language-agnostic data serialization format used for defining the service and message types.

### Repository Structure

- `pkg/db`: Handles database operations, including connecting to PostgreSQL and performing CRUD operations on the products table.
- `pkg/models`: Defines the data models used in the application, such as the `Product` model.
- `pkg/pb`: Contains the protocol buffer definition file (`product.proto`) and the generated Go code for the service and message types.
- `pkg/services`: Implements the server-side logic for the product service, including creating products, finding products, and managing product stock.
- `cmd/product-svc/main.go`: The main entry point of the application, initializes dependencies, and starts the gRPC server.


### Contributions

Contributions are welcome! If you'd like to contribute to the development of this microservice, please [fork the repository] and submit your pull requests.

### License
none

### Contact

For any inquiries or questions, feel free to contact the project maintainer at [rohither12@outlook.com]

