## To bring up the project run:

1. To bring up database on docker in root folder run: docker-compose up -d
2. In app directory run: go run main.go OR go build -o main, ./main

.env file contains DB connection variables

## Endpoints:

1. http://localhost:8081/customers (GET) - get all customers
2. http://localhost:8081/customer/{customerId} (GET) - get customer by id
3. http://localhost:8081/customers/top (GET) - get top customers
4. http://localhost:8081/customers/stats (GET) - get customers with a total number of each eaten  snack