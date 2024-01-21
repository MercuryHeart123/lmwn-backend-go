
# LINE MAN Wongnai Junior 2024.

This project is a JSON API built using Go, Go modules, and the Gin framework. The API provides COVID-19 summary data, counting the number of cases by provinces and age groups. The endpoint for accessing this data is /covid/summary.


## Requirements

- Go (version 1.21.5)
- Go modules
- Gin framework



## Installation

1.Clone the repository:

```bash
  git clone https://github.com/mercuryheart123/lmwn-backend-go.git
```
2.Navigate to the project directory:

```bash
  cd lmwn-backend-go
```
3.Install dependencies:

```bash
  go mod download
```
4.Run the application:
```bash
  go run main.go
```
The application should now be running on http://localhost:8080.


## API Endpoint

#### Get summarize case

Returns COVID-19 summary data, counting the number of cases by provinces and age groups.
```
  GET /covid/summary
```

#### Response
The response is a JSON object with the following structure:


```
{
  "Provinces": {
    "Amnat Charoen": 17,
    "Ang Thong": 36,
    "Bangkok": 27,
    "Bueng Kan": 23,
    "Buriram": 18,
    "Chachoengsao": 24,
    "Chai Nat": 25,
    // ...
  },
  "AgeGroup": {
    "0-30": 25,
    "31-60": 30,
    "60+": 12,
    "N/A": 5
  }
}

```
- The `Provinces` object contains data for each province, with case counts.
- The `AgeGroup` object contains data for all province, with age counts.
## Running Tests

To run tests, run the following command

```bash
  go test ./covid
```

