
## Run Locally

Clone the project

```bash
  git clone https://github.com/Kenasvarghese/Booking-App.git
```

Go to the project directory

```bash
  cd Booking-App/Backend
```

Run postgres instance

```bash
  docker compose up -d
```

Run migration

```bash
  goose up
```

Start the server

```bash
  go run main.go
```

     