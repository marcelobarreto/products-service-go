# Products Service with Go lang

Small CRUD application with Go, using Fiber + GORM.

## GET /products

List all products

### Response:

```json
[
  {
    "id": 1,
    "code": "USD187",
    "name": "",
    "price": 43.9,
    "discount": 10,
    "store_id": 1,
    "store": {
      "id": 1,
      "name": "Foo Store",
      "Open": true
    }
  },
  {
    "id": 2,
    "code": "USD187",
    "name": "",
    "price": 43.9,
    "discount": 10,
    "store_id": 1,
    "store": {
      "id": 1,
      "name": "Foo Store",
      "Open": true
    }
  },
  {
    "id": 3,
    "code": "USD187",
    "name": "",
    "price": 43.9,
    "discount": 10,
    "store_id": 1,
    "store": {
      "id": 1,
      "name": "Foo Store",
      "Open": true
    }
  },
  {
    "id": 5,
    "code": "BRL1187",
    "name": "",
    "price": 43.9,
    "discount": 10,
    "store_id": 1,
    "store": {
      "id": 1,
      "name": "Foo Store",
      "Open": true
    }
  },
  {
    "id": 6,
    "code": "FOO187",
    "name": "",
    "price": 43.9,
    "discount": 10,
    "store_id": 1,
    "store": {
      "id": 1,
      "name": "Foo Store",
      "Open": true
    }
  },
  {
    "id": 7,
    "code": "FOO187",
    "name": "",
    "price": 43.9,
    "discount": 10,
    "store_id": 1,
    "store": {
      "id": 1,
      "name": "Foo Store",
      "Open": true
    }
  },
  {
    "id": 8,
    "code": "FOO187",
    "name": "",
    "price": 43.9,
    "discount": 10,
    "store_id": 1,
    "store": {
      "id": 1,
      "name": "Foo Store",
      "Open": true
    }
  },
  {
    "id": 2002,
    "code": "BRL1187",
    "name": "",
    "price": 43.9,
    "discount": 10,
    "store_id": 1,
    "store": {
      "id": 1,
      "name": "Foo Store",
      "Open": true
    }
  }
]
```

## GET /products/:id

List specific product

### Response

```json
{
  "id": 1,
  "code": "USD187",
  "name": "",
  "price": 43.9,
  "discount": 10,
  "store_id": 1,
  "store": {
    "id": 1,
    "name": "Foo Store",
    "Open": true
  }
}
```

## POST /products

Creates a new product

### Body

```json
{
	"price": 43.90,
	"code": "FOO187",
	"discount": 10.0,
	"store_id": 1
}
```

### Response

```json
{
  "id": 8,
  "code": "FOO187",
  "name": "",
  "price": 43.9,
  "discount": 10,
  "store_id": 1,
  "store": {
    "id": 1,
    "name": "Foo Store",
    "Open": true
  }
}
```

## PUT /products/:id

Updates a product

### Body

```json
{
	"price": 43.90,
	"code": "BRL1187",
	"discount": 10.0
}
```

### Response

```json
{
  "id": 5,
  "code": "BRL1187",
  "name": "",
  "price": 43.9,
  "discount": 10,
  "store_id": 1,
  "store": {
    "id": 1,
    "name": "Foo Store",
    "Open": true
  }
}
```

## DELETE /products/:id

Delete a product

### Response

Status: OK