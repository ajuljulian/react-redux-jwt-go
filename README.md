# Description

A React JS app with a Go backend.

The React app provides the ability to register for an account, sign in, and access both protected and unprotected resources (based on user roles). It is based on the following example code I found:

https://www.bezkoder.com/react-hooks-redux-login-registration-example/

The Go backend exposes services for signing up for an account, signing in (using JWT), and accessing both protected and unprotected resources.

The entire system can be deployed using docker-compose.

# API

The api is written in Go.  It uses the [Echo](https://echo.labstack.com) framework for web services, [GORM](https://gorm.io) for ORM, and [Go JWT](https://github.com/golang-jwt/jwt) for issuing/checking JWT tokens for access.

```
go get github.com/labstack/echo/v4
go get github.com/golang-jwt/jwt
go get -u gorm.io/gorm
go get gorm.io/driver/postgres
```

# Web

I found [this excellent tutorial](https://www.bezkoder.com/react-hooks-redux-login-registration-example/) on React and Redux by [BezKoder](https://www.bezkoder.com/author/bezkoder/) and I'm using it for the web portion.  I've only had to make a few changes so far:

1. Upgrade the dependencies (such as react router)
1. Use bootstrap 5
1. Integrate with my own Go api

# Docker

You can deploy everything using docker-compose:
```
$ docker-compose up --build
```

Remove images:
```
$ docker-compose down --rmi all 
```

It deploys several containers:
1. api - exposes backend services
1. web - the react js app
1. postgres - the database
1. nginx - for routing

Once all the containers are up, you can access the app from:

http://localhost:3050/





