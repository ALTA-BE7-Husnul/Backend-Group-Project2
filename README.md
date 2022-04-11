
# E-Commerce

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
