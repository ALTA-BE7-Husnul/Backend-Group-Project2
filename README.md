
# E-Commerce

## About The Project
The simple E-commerce App with following MVP (Minumum Viable Product) :
- Users can operate create, read, update, and delete on user's data.
  - Besides, created user has access to login.
- Users can see the list of the products (CRUD).
  - Not only to see, users also can create, update, dan delete their product (only their product)
  - User can see all the products (their own products or others').
- Users can add products to shopping cart (CRUD).
  - Users can see, edit, and delete products from their own shopping cart.
- Users can order the products contained in their shopping cart.
  - Users can see history order and cancel order.
## API Reference


#### Add New Users 
```http
  POST  https://satriacening.cloud.okteto.net/users
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
|    request body   | `string` | **Required**. data to be inputted through the request body|

####  Login
```http
  POST https://satriacening.cloud.okteto.net/login
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `token`      | `string` | **Required**. tokens for login |



#### Get Users by id

```http
  GET  https://satriacening.cloud.okteto.net/users/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `token`   | `string` | **Required**. login |
| `id`      | `string` | **Required**. Id of user to fetch |


#### DELETE Users 
```http
  DELETE  https://satriacening.cloud.okteto.net/users/:id
```
only user that are entered according to the user

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `token`   | `string` | **Required**. login |
|    `id`   | `int` | **Required**. id of user to delete|

#### UPDATE Users 
```http
  PUT  https://satriacening.cloud.okteto.net//users/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `token`   | `string` | **Required**. login |
|    `id`   | `int` | **Required**. id of user to UPDATE followed by data via request body|

Click [here](https://app.swaggerhub.com/apis-docs/husnulnawafil27/ecommerce_app/1.0.0#/) to see more Open API Resource. 
