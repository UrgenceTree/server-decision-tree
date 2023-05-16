# Service Configuration Documentation

This is the documentation for the configuration of our Golang service, which uses RabbitMQ as a message broker and has a User API for handling user-related operations.

## Configuration File

The service configuration is defined in a JSON file. The location of this file is specified by the tree_config_file key in the configuration object. The value of tree_config_file should be a string representing the path to the JSON configuration file.

In the provided example, the configuration file is named tree_conf.json and is expected to be in the same directory as the service. If the configuration file is located elsewhere, adjust the path accordingly.

## RabbitMQ

The rabbitmq object in the configuration file specifies the settings for connecting to RabbitMQ.

- **uri**: This is the connection string for RabbitMQ. It includes the username, password, hostname, and virtual host for RabbitMQ.
- **queueName**: This is the name of the RabbitMQ queue that the service will interact with.
- **port**: This is the port number on which RabbitMQ is running.

Please replace "amqp://guest:guest@localhost:5672/" and "user_queue" with your actual RabbitMQ connection string and queue name.

## User API
The userAPI object in the configuration file specifies the settings for the User API.

- **baseUrl**: This is the base URL of the User API. All endpoint paths will be appended to this base URL.
endpoints: This is an object that defines the endpoint paths for various user-related operations. Each key is the name of an operation, and each value is the path for that operation.

In the endpoints object:

- **getUser**: The endpoint for getting information about a user. The {id} in the path should be replaced with the ID of the desired user.
- **createUser**: The endpoint for creating a new user.
- **updateUser**: The endpoint for updating a user's information. The {id} in the path should be replaced with the ID of the user to be updated.
- **deleteUser**: The endpoint for deleting a user. The {id} in the path should be replaced with the ID of the user to be deleted.

Please replace "http://localhost:8080/api" and the endpoint paths with the actual base URL and endpoints of your User API.

### Example

Here's an example of a complete configuration:

```json
{
    "tree_config_file": "tree_conf.json",
    "rabbitmq": {
        "uri": "amqp://guest:guest@localhost:5672/",
        "queueName": "user_queue",
        "port": 5672
    },
    "userAPI": {
        "baseUrl": "http://localhost:8080/api",
        "endpoints": {
            "getUser": "/users/{id}",
            "createUser": "/users",
            "updateUser": "/users/{id}",
            "deleteUser": "/users/{id}"
        }
    }
}
```

Please replace the example values with your actual settings before using this configuration.
