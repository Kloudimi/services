---
title: users
servicename: users
labels: 
- Readme
- Backend
---
Usage management and authentication

# Users Service

The user service provides user management, storage and authentication.

## cURL


### Users Create
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/users/Users/Create' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "email": "string",
  "first_name": "string",
  "last_name": "string",
  "password": "string"
};
# Response
{
  "token": "string",
  "user": {
    "email": "string",
    "first_name": "string",
    "id": "string",
    "last_name": "string"
  }
}
```


### Users Delete
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/users/Users/Delete' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "id": "string"
};
# Response
{}
```


### Users List
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/users/Users/List' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {};
# Response
{
  "users": [
    {
      "email": "string",
      "first_name": "string",
      "id": "string",
      "last_name": "string"
    }
  ]
}
```


### Users Login
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/users/Users/Login' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "email": "string",
  "password": "string"
};
# Response
{
  "token": "string",
  "user": {
    "email": "string",
    "first_name": "string",
    "id": "string",
    "last_name": "string"
  }
}
```


### Users Logout
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/users/Users/Logout' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "id": "string"
};
# Response
{}
```


### Users Read
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/users/Users/Read' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "ids": [
    "string"
  ]
};
# Response
{
  "users": [
    {
      "key": "string",
      "value": {
        "email": "string",
        "first_name": "string",
        "id": "string",
        "last_name": "string"
      }
    }
  ]
}
```


### Users ReadByEmail
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/users/Users/ReadByEmail' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "emails": [
    "string"
  ]
};
# Response
{
  "users": [
    {
      "key": "string",
      "value": {
        "email": "string",
        "first_name": "string",
        "id": "string",
        "last_name": "string"
      }
    }
  ]
}
```


### Users Update
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/users/Users/Update' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "email": {},
  "first_name": {},
  "id": "string",
  "last_name": {},
  "password": {}
};
# Response
{
  "user": {
    "email": "string",
    "first_name": "string",
    "id": "string",
    "last_name": "string"
  }
}
```


### Users Validate
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/users/Users/Validate' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "token": "string"
};
# Response
{
  "user": {
    "email": "string",
    "first_name": "string",
    "id": "string",
    "last_name": "string"
  }
}
```


