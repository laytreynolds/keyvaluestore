# kv_store

## Description
The Key Value HTTP Store is a simple in-memory key-value store that exposes a set of HTTP endpoints for storing, retrieving, and managing key-value pairs. This project allows users to interact with the store via a RESTful API, making it easy to integrate into other applications or services.

## Usage

### Installation

To get started with the Key Value HTTP Store, follow these steps:

**Clone the repository**:

### Build the project:
Run the following command to build the project:
`bash go build -o kvstore`

### Access the API:
Once the server is running, you can access the API at `http://localhost:8080`. You can use tools like `curl`, Postman, or your web browser to interact with the endpoints.


### Endpoints
Once the server is running (localhost:8080), you can interact with the key-value store using the following HTTP endpoints:

### Ping
- **URL**: `kvs/ping`
- **Method**: `GET`
- **Description**: A simple endpoint to check if the server is running.

### GetRequest
- **URL**: `kvs/get?key=<your_key>`
- **Method**: `GET`
- **Description**: Retrieve the value associated with the specified key.

### Add
- **URL**: `kvs/add?key=<your_key>`
- **Method**: `POST`
- **Body**: `{"<your_value>"}`
- **Description**: Add a new key-value pair. Fails if the key already exists.

### GetRequest All
- **URL**: `kvs/getall`
- **Method**: `GET`
- **Description**: Retrieve all key-value pairs in the store.

### Exists
- **URL**: `kvs/exists?key=<your_key>`
- **Method**: `GET`
- **Description**: Check if the specified key exists in the store.

### Count
- **URL**: `kvs/count`
- **Method**: `GET`
- **Description**: GetRequest the total number of key-value pairs in the store.

### Clear
- **URL**: `kvs/clear`
- **Method**: `POST`
- **Description**: Clear all key-value pairs from the store.

### Delete
- **URL**: `kvs/delete?key=<your_key>`
- **Method**: `DELETE`
- **Description**: Delete the specified key from the store.

### Update
- **URL**: `kvs/update?key=<your_key>`
- **Method**: `PUT`
- **Body**: `{"<your_value>"}`
- **Description**: Update the value of an existing key. Fails if the key does not exist.

### Upsert
- **URL**: `kvs/upsert?key=<your_key>`
- **Method**: `POST`
- **Body**: `{"<your_value>"}`
- **Description**: Insert a new key-value pair or update the existing key with a new value. 

## Server Configuration

The server listens on port `8080` by default. You can change the port by modifying the `PORT` constant in the code.

### Graceful Shutdown

The server supports graceful shutdown, allowing it to complete ongoing requests before shutting down. You can stop the server by sending an interrupt signal (e.g., `Ctrl+C`).

### Future Improvements

Adding worker pools to increase transaction speed and allow for multiple workers to handle requests from a single channel. This would replace the current for select pipeline pattern.
