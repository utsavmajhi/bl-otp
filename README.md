# (Bookollab OTP service)bl-otp

A OTP server for verifying user


## Features
    - Configurable OTP mail layout
    - gRPC server

## Start server

### Start mysql server
```
docker-compose up -d datastore
```

#### Create Database and Table
```
docker exec -it otp_datastore bash
mysql -p
CREATE DATABASE bookollab;
USE bookollab;
CREATE TABLE otp(
	email VARCHAR(200) PRIMARY KEY,
	otp VARCHAR(4) NOT NULL,
	expiry_time BIGINT NOT NULL
);
```
### Start OTP server

#### Create `.env` file with fields
```
COMPOSE_PROJECT_NAME=bookollab
SMTP_HOST=("smtp.gmail.com" for gmail)
SMTP_PORT=(587 for gmail)
EMAIL_FROM=(sender's email)
PASSWORD=(password of sender's email)
```
#### start docker container
```
docker-compose up otp_server
```

## Dependencies
1. server definition (proto files) https://github.com/Zzocker/bl-proto 
2. Utils https://github.com/Zzocker/bl-utils