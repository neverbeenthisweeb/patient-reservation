# API Contract

## Create Reservation

POST /reservations

### HTTP 200 OK

```
curl -X POST \
  'localhost:4040/reservations' \
  --header 'Accept: */*' \
  --header 'User-Agent: Thunder Client (https://www.thunderclient.com)' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "patient_id": 2,
  "doctor_id": 2,
  "slot_id": 1
}'

{
  "id": 3,
  "patient_id": 2,
  "doctor_id": 2,
  "started_at": "2023-04-09T10:00:00Z",
  "ended_at": "2023-04-09T12:00:00Z",
  "is_cancelled": false,
  "created_at": "2023-04-10T01:33:25.031432145+07:00",
  "updated_at": "2023-04-10T01:33:25.031435008+07:00"
}
```

## Cancel Reservation

PUT /reservations

### HTTP 200 OK

```
curl -X PUT \
  'localhost:4040/reservations' \
  --header 'Accept: */*' \
  --header 'User-Agent: Thunder Client (https://www.thunderclient.com)' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "reservation_id": 3
}'

{
  "id": 1,
  "patient_id": 1,
  "doctor_id": 1,
  "started_at": "2023-04-10T10:15:00+07:00",
  "ended_at": "2023-04-10T10:45:00+07:00",
  "is_cancelled": true,
  "created_at": "2023-04-10T02:00:00+07:00",
  "updated_at": "2023-04-10T02:00:00+07:00"
}
```

## Get List of Reservations

GET /reservations?show_cancelled=true

### HTTP 200 OK

```
curl -X GET \
  'localhost:4040/reservations' \
  --header 'Accept: */*' \
  --header 'User-Agent: Thunder Client (https://www.thunderclient.com)'

[
  {
    "id": 1,
    "patient_id": 1,
    "doctor_id": 1,
    "started_at": "2023-04-10T10:15:00+07:00",
    "ended_at": "2023-04-10T10:45:00+07:00",
    "is_cancelled": false,
    "created_at": "2023-04-10T02:00:00+07:00",
    "updated_at": "2023-04-10T02:00:00+07:00"
  },
  {
    "id": 2,
    "patient_id": 2,
    "doctor_id": 1,
    "started_at": "2023-04-10T11:30:00+07:00",
    "ended_at": "2023-04-10T13:30:00+07:00",
    "is_cancelled": true,
    "created_at": "2023-04-10T02:00:00+07:00",
    "updated_at": "2023-04-10T02:00:00+07:00"
  }
]
```

## Get Slots

GET /reservations/slots

### HTTP 200 OK

```
curl -X GET \
  'localhost:4040/reservations/slots' \
  --header 'Accept: */*' \
  --header 'User-Agent: Thunder Client (https://www.thunderclient.com)'

[
  {
    "id": 1,
    "started_at": "10:00",
    "ended_at": "12:00",
    "created_at": "2023-04-10T02:00:00+07:00",
    "updated_at": "2023-04-10T02:00:00+07:00"
  },
  {
    "id": 2,
    "started_at": "12:00",
    "ended_at": "14:00",
    "created_at": "2023-04-10T02:00:00+07:00",
    "updated_at": "2023-04-10T02:00:00+07:00"
  },
  {
    "id": 3,
    "started_at": "14:00",
    "ended_at": "16:00",
    "created_at": "2023-04-10T02:00:00+07:00",
    "updated_at": "2023-04-10T02:00:00+07:00"
  }
]
```
