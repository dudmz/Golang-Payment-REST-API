# Stone Pagamentos Challenge - Payment and Transfer REST API in Go Programming Language

This is a challenge set by Stone Pagamentos and solved by Carlos Damázio (github.com/carlosdamazio).

## Explanation
I was asked to create a payment API REST in Go programming language, that obeys some of it's constraints
regarding code delivery, code quality, documentation and clean code. In this implementation,
an Account and Transfer representation must be specified in order to create a payment environment.

## Project structure

This project consists of a REST API with a Dockerfile and docker-compose.yml file to make the
environment easier to set up. It also uses a MongoDB, a NoSQL database to store all Accounts and Transfer
entities.

## How to run project

```
$ docker-compose up -d --build

# To stop the project...

$ docker-compose down
```


## API REST Specification

Rule of thumb: JSON format only.

## Account entity

### Attributes

- id
- name
- cpf
- balance
- created_at

### Methods

#### GET /accounts/

- Request body: None;
- Response: List of Accounts.

#### GET /accounts/{account_id}/balance

- Request body: None;
- Response: Account balance.

#### POST /accounts/

- Request body: "name" and "cpf", all required;
- Response: CREATED if serialization is OK or BAD REQUEST if not.

### Business Constraints

- Balance might be initialized with any value to make it simple (default is 300).

## Transfer entity

### Attributes

- id 
- account_origin_id
- account_destination_id
- amount
- created_at

### Methods

#### GET /transfers/

- Request body: None;
- Response: List of Transfers.

#### POST /transfers/

- Request body: "originAccount", "destAccount" and "amount", all required;
- Response: CREATED if serialization is OK and follows business constraints or BAD REQUEST if not.

### Business constraints

- Must update both accounts balances presented in a transfer;
- If the transfer requester (origin account) doesn't have enough funds, must return an error code. 
