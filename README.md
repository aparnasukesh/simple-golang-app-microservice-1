# Golang Microservices Application with Docker and Kubernetes

This project demonstrates a simple implementation of two Golang microservices containerized with Docker and deployed on a Kubernetes cluster. The microservices interact with PostgreSQL and Redis for storage and caching. They also communicate using gRPC.

---

## **Features**

### **Microservice 1**
1. **Create a New User**
   - Accepts user details (`name`, `email`, etc.) as input.
   - Stores the user information in a PostgreSQL database.
   - Caches the user details in Redis for faster retrieval.

2. **Get a User by ID**
   - Accepts a user ID as input.
   - Checks the Redis cache for the user details.
   - If found in cache, returns the details.
   - If not found, retrieves the details from the PostgreSQL database and caches them.
   - Returns an appropriate response if the user does not exist.

3. **Update a User**
   - Accepts a user ID and updated details as input.
   - Updates the user information in the PostgreSQL database.
   - Updates the Redis cache if the user exists.

4. **Delete a User**
   - Accepts a user ID as input.
   - Deletes the user from the PostgreSQL database.
   - Removes the user details from the Redis cache if they exist.

---

### **Microservice 2**
1. **Method Execution Endpoint**
   - **Method 1** (Sequential Execution)
     - Processes requests sequentially.
     - Checks the number of users in the database and their names.
     - Waits for the specified `waitTime` (in seconds) before returning the user list.
   - **Method 2** (Parallel Execution)
     - Processes requests in parallel.
     - Checks the number of users in the database and their names.
     - Waits for the specified `waitTime` (in seconds) before returning the user list.

2. **Communication with Microservice 1**
   - Uses gRPC to interact with Microservice 1 for user data retrieval.

---

## **Technologies Used**
- **Programming Language**: Golang
- **Database**: PostgreSQL
- **Cache**: Redis
- **Communication**: gRPC
- **Containerization**: Docker
- **Orchestration**: Kubernetes

---

## **Setup and Deployment**

### **Prerequisites**
- Docker installed
- Kubernetes cluster (local or cloud-based, e.g., Minikube or Kubernetes on Docker Desktop)
- Helm (optional, for easier Kubernetes deployments)

---

### **Step 1: Clone the Repository**
```bash
git clone https://github.com/aparnasukesh/simple-golang-app-microservice-1.git
git clone https://github.com/aparnasukesh/simple-golang-app-microservice-two.git
cd yourproject
