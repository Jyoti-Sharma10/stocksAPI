# stocksAPI

Stocks API
This is a simple API built with Go, Gorilla Mux, and PostgreSQL for performing operations on stock data

API Endpoints
Get all stocks
Method: GET
Endpoint: /api/stock
Response: JSON array of stock objects

Get a single stock
Method: GET
Endpoint: api/stock/{id}
Response: JSON object representing a stock

Create a new stock
Method: POST
Endpoint: api/newstock
Request Body: JSON object representing the stock to be created
Response: JSON object representing the created stock

Update a stock
Method: PUT
Endpoint: api/stock/{id}
Request Body: JSON object representing the updated stock data
Response: JSON object representing the updated stock

Delete a stock
Method: DELETE
Endpoint: api/deletestock/{id}
Response: JSON object containing a success message
Please note that {id} should be replaced with the ID of the stock you want to perform the respective operation on.

Feel free to customize this README file based on your specific project needs and requirements. 







