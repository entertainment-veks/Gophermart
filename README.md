# Gophermart
Demo project which implements accumulative loyalty system

### Endpoints
#### User registration
```
POST /api/user/register
```
Registration is carried out using a pair of login / password. Each login is unique.  
If registration is successful, the user is automatically authenticated.

##### Request format:
```
POST /api/user/register HTTP/1.1
Content-Type: application/json
...

{
    "login": "<login>",
    "password": "<password>"
} 
```
##### Possible response codes:
* 200 - the user is successfully registered and authenticated;
* 400 - invalid request format;
* 409 - login is already taken;
* 500 - internal server error.

#### User Authentication
```
POST /api/user/login
```
Аутентификация производится по паре логин/пароль.

##### Request format:
```
POST /api/user/login HTTP/1.1
Content-Type: application/json
...

{
    "login": "<login>",
    "password": "<password>"
}
```
##### Possible response codes:
* 200 - the user is successfully authenticated;
* 400 - invalid request format;
* 401 - invalid login / password pair;
* 500 - internal server error.

#### Load order number
```
POST /api/user/orders
```
The handler is available only to authenticated users.  
The order number is a sequence of digits of arbitrary length.  
The order number must match the Luna algorithm.

##### Request format:
```
POST /api/user/orders HTTP/1.1
Content-Type: text/plain
...

12345678903
```
##### Possible response codes:
* 200 - the order number has already been uploaded by this user;
* 202 - the new order number is accepted for processing;
* 400 - invalid request format;
* 401 - the user is not authenticated;
* 409 - the order number has already been uploaded by another user;
* 422 - invalid order number format;
* 500 - internal server error.

#### Getting a list of uploaded order numbers
```
GET /api/user/orders
```
The handler is available only to an authorized user.  
Order numbers are sorted by load time from oldest to newest.  
The date format is RFC3339.

##### Available statuses of processing calculations:
```
* NEW - the order has been uploaded to the system, but not processed;
* PROCESSING - the reward for the order is calculated;
* INVALID - the reward calculation system refused to pay;
* PROCESSED - the order data has been verified and  
the payment information has been successfully received.
```

##### Request format:
```
GET /api/user/orders HTTP/1.1
Content-Length: 0 
```
##### Possible response codes:
* 200 - successful request processing;
* 204 - no data to answer;
* 401 - the user is not authorized;
* 500 - internal server error.

##### For start
Start migrations: 
> migrate -path migrations -database "postgres://localhost/gophermart?sslmode=disable" up
