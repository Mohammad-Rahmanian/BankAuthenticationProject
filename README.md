# Bank Authentication Service Project

This project is implemented in the Golang programming language and is designed to create a bank authentication service. The primary objective of this project is to work with various cloud services, including MongoDB Atlas for database storage, an S3-compatible object storage service provided by Abararvan for storing photos, CloudAMQP for RabbitMQ, IMGGA for image processing, and Mailgun for email sending.

## Project Description

The software consists of two backend services:

### First Service

The first service is responsible for receiving user requests and responding to them. It includes two APIs:

**1. Request Registration API:**
- This API receives user information, which includes email, last name, national code, IP address, and two photos.
- It stores the user's information in the database and generates a unique username.
- The API also stores the two individual images in an object storage.
- Afterward, the person's username is placed in the RabbitMQ queue.
- As a response to the request, a message like "Your authentication request was registered" is displayed to the user.

**2. Request Status Check API:**
- This API receives a person's national code.
- If the request is queued, it returns the message "in review."
- If the request is rejected, it returns the message "Your authentication request has been rejected. Please try again later."
- If the request is approved, it returns the message "Authentication was successfully completed."

### Second Service

The second service is responsible for processing messages from the RabbitMQ queue and storing the results in the database. Here's how it works:

- This service is connected to the RabbitMQ queue and listens for new messages.
- Each message read from the queue contains a username. The service retrieves the person's photos from the object store based on the username.
- Both photos are sent separately to the facial recognition service. If there is no face detected, the person's request is rejected.
- If a face is detected, the service uses the facial similarity service to compare the two photos using the identifiers received from the previous step. If the similarity is more than 80%, the person's request is accepted, and the status in the database is updated.
- Using the email sending service (Mailgun), an email is sent to the user to inform them of their authentication status.

## Configuration

**Setup:**

1. **Create Config File:**
   - Navigate to the project directory.
   - Create a `configs` directory.
   - Inside `configs`, create a file named `config.go` with the following template (replace placeholder values with your actual API keys and endpoints):

```go
package configs

var (
    MailApiKey               = "<MAIL_API_KEY>"
    MailDomain               = "<MAILGUN_DOMAIN>"
    MailSender               = "<EMAIL_SENDER>"
    ImageProcessingApiKey    = "<IMG_PROCESSING_API_KEY>"
    ImageProcessingSecretKey = "<IMG_PROCESSING_SECRET_KEY>"
    MessageBrokerURL         = "<RABBITMQ_URL>"
    DatabaseURL              = "<MONGODB_URL>"
    StorageServiceID         = "<STORAGE_SERVICE_ID>"
    StorageServiceSecret     = "<STORAGE_SERVICE_SECRET>"
    StorageServiceEndpoint   = "<STORAGE_ENDPOINT>"
)
```
## Run the Services:
Start the services in the specified order:

Run go run ./firstService/main.go to start the First Service.

After the First Service is running, execute go run ./secondService/main.go to start the Second Service.
  
**Note:** Ensure that you have the necessary API keys and access to the specified cloud services before running the project.
