# Booking API
Booking API is a simple API for reservation management in a boat with different staterooms. Supports the next features:

-  Booking the full boat for N days
-  Booking just N staterooms of the boat
-  Booking the boat in shared way. This mean, other client's can join to your reservation.


# API usage
This section won't get in details deeply (explanation of the parameters, possible returns, etc) because it's not the porpouse of the repository, but the explanation of the Software engeneering areas involved in its construction.

The application counts with the next endpoints:
### `POST /auth/register`
Creates a new user, expects a JSON body request like this:
```json
{
  "email": "myEmail@gmail.com",
  "password":"secretPassword"
}
```
### `POST /auth/login`
Authenticate the user and retrieve the Json Web Token, expects a JSON body request like this:
```json
{
  "email": "myEmail@gmail.com",
  "password":"secretPassword"
}
```
### `POST /boat`
Creates a new boat, expects a JSON body request like this:
```json
{
  "name": "BoatName",
  "maxCapacity":10
}
```
| Parameter | Type   | Description                            |
| --------------------- | ------ | --------------------------------------|
| `name`                | String | Boat's name                  |
| `maxCapacity`         | Number | Boat's max capacity      |

If successful it will return a JSON like the following

```json
{
    "id": 14,
    "maxCapacity": 10,
    "name": "BoatName",
    "owner": "myEmail@gmail.com",
    "stateRoom": null
}
```
### `PUT /stateRoom/add/:boatId`
Adds a new stateroom to a boat, given the boat ID

### `GET /boat/:id`
Get a specific boat, given its ID. As the Create-Boat endpoint, returns a JSON like the following
```json
{
    "id": 14,
    "maxCapacity": 10,
    "name": "BoatName",
    "owner": "myEmail@gmail.com",
    "stateRoom": null
}
```

### `GET /notAvailableForShared/:boatId?passengers=:passengers`
Get the days when a boat it's in not avaialble for shared reservations, given the boat ID.
A day is not available for a shared resrevation if:

- Exists a not-shared reservation in those days
- Exist other shared reservation in those days. But, the sum of the passengers exceeds the maximum capacity.

The query-parameter `passengers` is the amount of passengers to check disponibility. The default value is 1.

Expects to return a JSON like the following:
```
{
    "days": [
        "2023-12-20",
        "2023-12-22",
        "2023-12-23",
        "2023-12-24"
    ]
}
```
### `GET /notAvailableForClose/:boatId?staterooms=:staterooms`
Get the days when a boat it's in not avaialble for not-shared (close) reservations, given the boat ID.
The query parameter `staterooms` are the needed staterooms for the reservation. The default value is 1.
A day is not available for a close resrevation if:

- Exists a shared reservation on those days
- There aren't enough staterooms

Expects to return a JSON like the following:
```
{
    "days": [
        "2023-12-20",
        "2023-12-22",
        "2023-12-23",
        "2023-12-24"
    ]
}
```

### `GET /boat/notEmpty/:id`
Get those days where there is at least one reservation of a boat given its ID

### `POST /boat/reservateFullBoat`
Makes a reservation in all the boat's stasterooms, expects a JSON body request like this:
```json
{
	"email":"email@gmail.com",
	"phone":"623029321",
	"firstDay":"2023-12-25",
	"lastDay":"2023-12-25",
	"boatId":1,
    	"isOpen":true,
    	"passengers":3
}
```

| Parameter | Type   | Description                            |
| --------------------- | ------ | --------------------------------------|
| `email`                | String | Client's email                |
| `phone`         | String | Client's phone number    |
| `firstDay`         | Date | Reservation's first day      |
| `lastDay`         | Date | Reservation's last day      |
| `boatId`         | Number | Boat's ID      |
| `isOpen`         | Boolean | If the reservation is shared or not      |
| `passengers`         | Number | Number of passengers if the reservation is shared      |

This endpoint could return an error if there is not enough space for the new reservation

### `POST /boat/reservate?staterooms=:staterooms`
Makes a reservation in the staterooms specified. 
The query parameter `staterooms` expects the staterooms's amount to reservate, the default value is 1.
Only not-shared reservations can reservate a specific amount of staterooms. Shared ones must reservate the full boat.
Expects a JSON body request like this:
```json
{
	"email":"email@gmail.com",
	"phone":"623029321",
	"firstDay":"2023-12-25",
	"lastDay":"2023-12-25",
	"boatId":1
}
```

| Parameter | Type   | Description                            |
| --------------------- | ------ | --------------------------------------|
| `email`                | String | Client's email                |
| `phone`         | String | Client's phone number    |
| `firstDay`         | Date | Reservation's first day      |
| `lastDay`         | Date | Reservation's last day      |
| `boatId`         | Number | Boat's ID      |

This endpoint could return an error if there is not enough space for the new reservation

### `DELETE /boat/:id`
Deletes a specific boat, given its ID

### `DELETE /reservation/:id`
Deletes a specific reservation, given its ID

## Reservations optimization
Sometimes, the reservations may need a reallocation to be able to insert a new one, which can't be allocated with the current reservations distribution.
For example, in the next reservation distribution, we would not be able to insert a new reservation between days 3 and 10, although this is possible if we reallocate all of them
![Reservation distribution](https://github.com/lucastomic/bookingAPI/assets/65186233/71ff2d40-895c-42f9-9576-a237f8b7f1ed)

The algorithm takes into account that the reservations which have already started, can't be reallocated (the clients could be already in their staterooms)

