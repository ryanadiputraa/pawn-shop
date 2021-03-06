# PAWN SHOP

Pawn Shop is a information system for managing Pawn Shop Information System data. Build with Golang as a RESTful API with PostgreSQL for database.

## API SPEC

---

## Employees

### Get All Employees

-   Method : `GET`
-   Endpoint : `/api/employees`
-   Header :
    -   Content-Type : `application/json`
    -   Accept : `application/json`
    -   X-Access-Token: `token`
-   Response :

```json
[
    {
        "employee_id": "Number",
        "firstname": "String",
        "lastname": "String",
        "gender": ("pria" || "wanita"),
        "birthdate": "YYYY-MM-DD",
        "address": "String",
        "password": "String"
    },
    {
        "employee_id": "Number",
        "firstname": "String",
        "lastname": "String",
        "gender": ("pria" || "wanita"),
        "birthdate": "YYYY-MM-DD",
        "address": "String",
        "password": "String"
    }
]
```

### Get Employee by Id

-   Method : `GET`
-   Endpoint : `/api/employees/{employee_id}`
-   Header :
    -   Content-Type : `application/json`
    -   Accept : `application/json`
    -   X-Access-Token: `token`
-   Response :

```json
{
    "employee_id": "Number",
    "firstname": "String",
    "lastname": "String",
    "gender": ("pria" || "wanita"),
    "birthdate": "YYYY-MM-DD",
    "address": "String",
    "password": "String"
}
```

### Get Employee by Name

-   Method : `GET`
-   Endpoint : `/api/employees/{employee_name}`
-   Header :
    -   Content-Type : `application/json`
    -   Accept : `application/json`
    -   X-Access-Token: `token`
-   Response :

```json
{
    "employee_id": "Number",
    "firstname": "String",
    "lastname": "String",
    "gender": ("pria" || "wanita"),
    "birthdate": "YYYY-MM-DD",
    "address": "String",
    "password": "String"
}
```

### Register Employee

-   Method : `POST`
-   Endpoint : `/api/employees`
-   Header :
    -   Content-Type : `application/json`
    -   Accept : `application/json`
    -   X-Access-Token: `token`
-   Body :

```json
{
    "firstname": "String",
    "lastname": "String",
    "gender": ("pria" || "wanita"),
    "birthdate": "YYYY-MM-DD",
    "address": "String",
    "password": "String"
}
```

-   Response :

```json
{
    "code": "Number"
}
```

### Update Employee

-   Method : `PUT`
-   Endpoint : `/api/employees/{employee_id}`
-   Header :
    -   Content-Type : `application/json`
    -   Accept : `application/json`
    -   X-Access-Token: `token`
-   Body :

```json
{
    "firstname": "String",
    "lastname": "String",
    "gender": ("pria" || "wanita"),
    "birthdate": "YYYY-MM-DD",
    "address": "String",
    "password": "String"
}
```

-   Response :

```json
{
    "code": "Number"
}
```

### Delete Employee

-   Method : `DELETE`
-   Endpoint : `/api/employees/{employee_id}`
-   Header :
    -   Content-Type : `application/json`
    -   Accept : `application/json`
    -   X-Access-Token: `token`
-   Response :

```json
{
    "code": "Number"
}
```

### Login Employee

-   Method : `POST`
-   Endpoint : `/auth/login`
-   Header :
    -   Content-Type : `application/json`
    -   Accept : `application/json`
-   Body :

```json
{
    "employee_id": "Number",
    "password": "String"
}
```

-   Response :

```json
{
    "code": "Number"
}
```

### Logout Employee

-   Method : `POST`
-   Endpoint : `/auth/logout`
-   Header :
    -   Content-Type : `application/json`
    -   Accept : `application/json`
-   Response :

```json
{
    "code": "Number"
}
```

---

## Customers

### Get All Customers

-   Method : `GET`
-   Endpoint : `/api/customers`
-   Header :
    -   Content-Type : `application/json`
    -   Accept : `application/json`
    -   X-Access-Token: `token`
-   Response :

```json
[
    {
            "customer_id": "String",
            "firstname": "String",
            "lastname": "String",
            "gender": ("pria" || "wanita"),
            "loan": "String",
            "insurance_item": "String",
            "contact": "String",
    },
    {
            "customer_id": "String",
            "firstname": "String",
            "lastname": "String",
            "gender": ("pria" || "wanita"),
            "loan": "String",
            "insurance_item": "String",
            "contact": "String",
    }
]
```

### Get Customer Data By Name

-   Method : `GET`
-   Endpoint : `/api/customers/{customer_name}`
-   Header :
    -   Content-Type : `application/json`
    -   Accept : `application/json`
    -   X-Access-Token: `token`
-   Response :

```json
[
    {
        "customer_id": "String",
        "firstname": "String",
        "lastname": "String",
        "gender": ("pria" || "wanita"),
        "loan": "String",
        "insurance_item": "String",
        "contact": "String",
    },
    {
        "customer_id": "String",
        "firstname": "String",
        "lastname": "String",
        "gender": ("pria" || "wanita"),
        "loan": "String",
        "insurance_item": "String",
        "contact": "String",
    },
]
```

### Add Loan

-   Method : `POST`
-   Endpoint : `/api/customers/{customer_id}`
-   Header :
    -   Content-Type : `application/json`
    -   Accept : `application/json`
    -   X-Access-Token: `token`
-   Body :

```json
{
    "customer_id": "String",
    "firstname": "String",
    "lastname": "String",
    "gender": ("pria" || "wanita"),
    "nominal": "Number",
    "interest": "Number",
    "insurance_item": "String",
    "contact": "String",
}
```

-   Response :

```json
{
    "code": "Number"
}
```

### Pay Off The Loan

-   Method : `PUT`
-   Endpoint : `/api/customers/{customer_id}`
-   Header :
    -   Content-Type : `application/json`
    -   Accept : `application/json`
    -   X-Access-Token: `token`
-   Response :

```json
{
    "code": "Number"
}
```

### Get Financial Statements

-   Method : `GET`
-   Endpoint : `/api/customers/financial`
-   Header :
    -   Content-Type : `application/json`
    -   Accept : `application/json`
    -   X-Access-Token: `token`
-   Response :

```json
{
    "totalLoans": "Number",
    "loanStatus": [
        {
            "status": "String",
            "total": "Number"
        },
        {
            "status": "String",
            "total": "Number"
        }
    ]
}
```
