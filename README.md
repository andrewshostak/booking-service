# booking-service

This is a golang microservice with a layered architecture. It has 3 endpoints:

- POST /bookings
    - validates the request payload
    - verifies if launchpad is available by calling [SpaceX-API](https://github.com/r-spacex/SpaceX-API)
    - creates a booking in a database
    - returns created data
- GET /bookings
    - returns an array of all database bookings 
- DELETE /bookings/:id
    - deletes a booking by id

### Run via docker-compose
`docker-compose up`

### Example requests

#### Create booking:
```
curl --location --request POST 'localhost:8080/bookings' \
--header 'Content-Type: application/json' \
--data-raw '{
    "first_name": "john",
    "last_name": "doe",
    "gender": "male",
    "birthday": "2012-02-11",
    "launchpad_id": "5e9e4502f509094188566f88",
    "destination_id": "mars",
    "launch_date": "2022-12-02"
}'
```
#### Create booking with unavailable launchpad:
```
curl --location --request POST 'localhost:8080/bookings' \
--header 'Content-Type: application/json' \
--data-raw '{
    "first_name": "john",
    "last_name": "doe",
    "gender": "male",
    "birthday": "2012-02-11",
    "launchpad_id": "5e9e4502f509094188566f88",
    "destination_id": "mars",
    "launch_date": "2022-12-01"
}'
```
#### List bookings
```curl --location --request GET 'localhost:8080/bookings'```
#### Delete booking
```curl --location --request DELETE 'localhost:8080/bookings/1'```

### Run tests

`go test ./...`