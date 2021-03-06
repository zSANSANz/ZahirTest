Rest Test.
• From Above ERD please create Rest full API.

![alt text](https://github.com/zSANSANz/ZahirTest/blob/master/erd.PNG)

1. Create register API(Include Generate password).
• Acceptance
o Phone number and email is unique.
o Customer name,email,phone number,dob,sex,created_date
is mandatory.
o Password generated using SHA256 etc mix with salt
key(dynamic).
• Negative case

2. Create get token api.
• Acceptance
o Phone_number_or_email and password is mandatory.
o Passed validation from (phone_number_or_email ) and
password.
o Must return token with access & refresh type

3. Create refresh token api.
• Acceptance.
o Must return token with access & refresh type

4. Create order api.
• Acceptance
o Passed validation from bearer auth.
o token is only one use.
o order number generated with format PO-123/IX/2020 (IX is
current month)(2020 is current year),(123 reset per month).
o order detail can be more than one
# MINI-Project e-Commerce REST API (Zahir)

## Overview

Zahir MINI-Project e-Commerce is REST-API specifically build to support online store system of Zahir. Zahir provides functionality that allows developers to behave either as customer or admin. Zahir is created using Golang, Gorm, and MYSQL as database.

The MVP of Zahir MINI-Project e-Commerce is REST-API is :

* Login and register customers
* Customers can view list product
* Customers can add product to order
* Customers can delete the product list from the order detail
* Customers can make payment transactions

And this ERD of Zahir MINI-Project e-Commerce is REST-API is like this :

![alt text](https://github.com/zSANSANz/ZahirTest/blob/master/erd.PNG)

## Tutorial

We provide another documentation using swagger to make it easy to understand the basic of API.

How to run this project :

1. git clone `git@github.com:zSANSANz/ZahirTest.git`
2. cd `ZahirTest` and run `go install`
3. setup envirotment, rename file `.env.example` to `.env`. And setup this variabel envirotment.
4. after finished install this module, run command `go run swagger.go` for running this documentation to browser like this `localhost:8080/swaggerui`
5. run command `go run main.go` to running this server like this `localhost:8000`
6. run command `go test ./controller/ -cover` for running unit testing

## HTTP requests


There are 4 basic HTTP requests that you can use in this API:

* `POST` Create a resource
* `PUT` Update a resource
* `GET` Get a resource or list of resources
* `DELETE` Delete a resource

## HTTP Responses

Each response will include a code(repsonse code),message,status and data object that can be single object or array depending on the query.

## HTTP Response Codes

Each response will be returned with one of the following HTTP status codes:

* `200` `OK` The request was successful
* `400` `Bad Request` There was a problem with the request (security, malformed, data validation, etc.)
* `404` `Not found` An attempt was made to access a resource that does not exist in the API
* `500` `Server Error` An error on the server occurred

## API Documentation

Retail Store Ecommerce API
- [Swagger Documentation](https://app.swaggerhub.com/apis/hasantech/simple-inventory_api/1.2)

### /register

#### POST
##### Summary

register customer

##### Description
passing customer data to register customer

### Parameter

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| username | body | customer usernmae | Yes | string
| email | body | customer email | Yes | string |
| password | body | customer password | Yes | string |

##### Responses
| Code | Description | 
| ---- | ----------- |
| 200 | success register customer |
| 500 | error registration customer | 

#### Request Body Parameter Example
```json
{
    "username":"lisa233",
    "email":"lisa@gmail5.com",
    "password":"312457"
}
```

#### Response Body Example
```json
{    
    "code": 200,
    "message": "success register customers",
    "status": "success"
}
```

### /login

#### POST
##### Summary

login as customer

##### Description
passing email and password to login as customer

####  Parameter
| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| email | body | customer email | Yes | string |
| password | body | customer password | Yes | string |

##### Responses
| Code | Description | 
| ---- | ----------- |
| 200 | success login customer |
| 400 | error login(invalid email/password) | 

#### Request Body Parameter Example
```json
{
   "email":"susi@gmail.com",
   "password":"john123"
}
```

#### Response Body Example
```json
{    
    "code": 200,
    "message": "success login customer",
    "status": "success",
    "data": {
        "id": 1,
        "email": "susi@gmail.com",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MTczNDg0ODgsInVzZXJJZCI6MX0.C4pvmawkwNP-SVJlIA0qbTct3qen4PPbhtHK9XRRpts"
    }
}
```
### /customers

### /customers/{id}

#### PUT
##### Summary

Update Customer Data

##### Description
passing several data to update  customer data 

####  Parameter
| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | customer id | Yes | string |
| email | body | customer email | No | string |
| adresss | body | customer address | No | string |
| bank_name | body | customer email | No | string |
| bank_account_number | body | customer address | No | string |

##### Responses
| Code | Description | 
| ---- | ----------- |
| 200 | success update customer data |
| 400 | error update data| 

#### Request Body Parameter Example
```json
{
    "email":"32ibu@gmail.com",
    "address": "jalan ahmad yani",
    "bank_name": "BCA",
    "bank_account_number": "23232323"
}
```

#### Response Body Example
```json
{    
    "code": 200,
    "message": "success update profil customer",
    "status": "success"
}
```

### /products

### /products?category={id}

##### GET
##### Summary

Get products

##### Description
Get all product based on category, if no id is passed then it will return all product in the store

#### Parameter
| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | query | category id | No | int |

##### Responses
| Code | Description | 
| ---- | ----------- |
| 200 | success get products |
| 400 | error bad request | 


#### Response Body Example
```json
{    
    "code": 200,
    "message": "success get product",
    "status": "success",
    "data":  [
        {
            "id": 1,
            "name": "Yonex Zerox",
            "categories_id": 4,
            "description": "Raket tercepat di Asia",
            "quantity": 30,
            "price": 45000,
            "unit": "pcs",
            "CreatedAt": "2021-03-28T20:01:40.21+07:00",
            "UpdatedAt": "2021-03-28T20:01:40.21+07:00"
        }
    ]
}
```

### /cartitems

#### POST
##### Summary

Adding cartitems to carts

##### Description
Adding product that customer wants to buy into their cart

####  Parameter
| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| carts_id | body | customer cart id | Yes | int |
| products_id | body | product id | Yes | int |
| quantity | body | quantity of product | Yes | int |

##### Responses
| Code | Description | 
| ---- | ----------- |
| 200 | success insert cartitems |
| 400 | invalid input| 
| 500 | product/cart not exist|

#### Request Body Parameter Example
```json
{
    "carts_id":4,
    "products_id":9,
    "quantity":3
}
```

#### Response Body Example
```json
{    
    "code": 200,
    "status": "success",
    "message": "success insert cartitems"
}
```

#### /cartitems?cart={cartid}

#### GET
##### Summary

get all cartitems inside specific cart

##### Description
get all cartitems that is inside usercart

####  Parameter
| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| carts_id | query | customer cart id | Yes | int |

##### Responses
| Code | Description | 
| ---- | ----------- |
| 200 | success get cartitems |
| 400 | bad request| 



#### Response Body Example
```json
{    
    "code": 200,
    "status": "success",
    "message": "Success Get Cartitems",
    "data": [
        {
            "id": 29,
            "carts_id": 4,
            "products_id": 9,
            "name": "Yonex e545",
            "price": 45000,
            "quantity": 3
        }
    ]
}
```

#### /cartitems/:cartitemid

#### DELETE
##### Summary

Delete specific cartitem

##### Description
Delete specific cartitem based on cartitemid

####  Parameter
| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| cartitemid | path |  id of cartitem to be deleted | Yes | int |

##### Responses
| Code | Description | 
| ---- | ----------- |
| 200 | success delete cartitems |
| 400 | error deleting cartitems(wrong cartitem id etc)| 



#### Response Body Example
```json
{    
     "code": 200,
     "message": "cartitems succesfully deleted",
     "status": "success"
}
```

### /orders

#### POST
##### Summary

Creating order from customer

##### Description
When customer checkout order and payment will be created, and customer cart item will be moved to checkout item

####  Parameter
| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| customers_id | body | customer id | Yes | int |
| courier_id | body | courier id | Yes | int |
| payment_method | body | method of payment | Yes | string |
| payment_start_date | body | date in format(yyyy-mm-dd hh:mm:ss)| Yes | string |
| payment_end_date | body | date in format(yyyy-mm-dd hh:mm:ss)| Yes | string |
| payment_status | body | status of payment | Yes | string |
| payment_amount | body | total payment amount | Yes | int |


##### Responses
| Code | Description | 
| ---- | ----------- |
| 200 | success create order |
| 400 | bad request(customer not exist,etc)| 

#### Request Body Parameter Example
```json
{
    "customers_id":4,
     "couriers_id":1,
     "payment_method":" BCA",
     "payment_start_date":"2021-03-29 10:42:44.710",
     "payment_end_date" : "2021-03-30 10:42:44.710",
     "payment_status": "waiting",
     "payment_amount" : 345000
}
```

#### Response Body Example
```json
{    
    "code": 200,
    "message": "success insert order",
    "status": "success"
}
```
### /payments/:id

#### PUT
##### Summary

Update payment status

##### Description
Update payment status when customer finishing their payment for the product

####  Parameter
| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | payment id that needs to be updated | Yes | int |

##### Responses
| Code | Description | 
| ---- | ----------- |
| 200 | success update payment |
| 400 | error update payment(payment id not exist etc)| 


#### Response Body Example
```json
{    
    "code": 200,
    "message": "success update payments",
    "status": "success"
}
```
