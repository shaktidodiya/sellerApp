## SellerApp

# Assumptions
1. I have assumed that already we have services for creating the order, manage inventory of products.
2. Assumed that middleware is already there for validation of trusted user.
3. Assumed that this service needs to take care for creating the order with UserID, ProductID, Price and Quantity. And to get the order or fetch the matching orders based on UserID, ProductID, Price and Quantity. Here UserID should be always available as it needs to get who logged in otherwise it will get another user orders also.
4. Assumed that product tables are also there and i have productId in order table refers to the product.

# How to run
To run the application follow the steps :
## Steps
    I'm assuming that the environment for database password and user are as i have used in the docker-compose and in the main.go.
    
1. Run `docker-compose up` this will install the required dependency docker image such as MySQL and GoLang.
2. this will start the application and it will run on PORT : 8000 and ready to serve
3. now you can use any RestAPI client to hit the endoint for Creating the order and Get the order based on matching. 
4. Use the attached Collection for hitting the endpoint.

## Steps to run without Docker
1.  First create the Database and DB table, use the `database.sql` file to create the table by copy pasting the query into MySQL terminal.
2. Now run `go run main.go` this will start the Application and it will run on port 8000. Change the HOST value as db to localhost in the main.go for if running mysql in local.
3. Use the RestAPI collection and hit the endpoint.
