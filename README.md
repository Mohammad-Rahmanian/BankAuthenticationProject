# Bank Authentication Service Project

This project is implemented in the Golang programming language and is designed to create a bank authentication service. The primary objective of this project is to work with various cloud services, including MongoDB Atlas for database storage, an S3-compatible object storage service provided by Abararvan for storing photos, CloudAMQP for RabbitMQ, IMGGA for image processing, and Mailgun for email sending.

## Project Description
file:///home/sajjad/Pictures/Screenshots/Screenshot%20from%202023-10-25%2017-41-59.png

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

To run this project, you need to place a configuration file in the `config` directory. This configuration file should contain the necessary API keys and URLs required to connect to the various cloud services used in the project.

Please ensure that the configuration file is properly populated with the required credentials and service endpoints before running the project.

For detailed instructions on how to set up and run the project, please refer to the project's documentation or user guide.

**Note:** Ensure that you have the necessary API keys and access to the specified cloud services before running the project.
