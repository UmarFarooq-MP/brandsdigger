# BrandsDigger

BrandsDigger is a REST service that takes an idea as input, generates potential brand names for that idea using OpenAI, and then checks the availability of these names as domain names via GoDaddy.

## Features
- **Brand Name Generation**: Uses OpenAI to generate creative names based on the provided idea.
- **Domain Availability Check**: Verifies the availability of generated names as domains using GoDaddy API.
- **Simple REST Interface**: Easy-to-use API for interacting with the service.

## Prerequisites
Before running the application, ensure you have the following environment variables set:

```sh
export GODADDY_API_KEY=your_godaddy_api_key
export GODADDY_API_SECRET=your_godaddy_api_secret
export OPENAI_API_KEY=your_openai_api_key
```

## Installation & Running
To run the service, simply execute:

```sh
go run cmd/main.go
```

## API Usage
### Generate Brand Names
#### Request
```sh
curl --location 'http://localhost:8080/generate/names' \
--header 'Content-Type: application/json' \
--data '{
    "id":"1",
    "message" : "I am a tomato ketchup company"
}'
```
#### Response (Example)
```json
["TomatoLuxe.com","TangyFuse.com","PureMato.com","KetchGrind.com","TomoChup.com","KetchaPop.com","SavoryDrip.com"]
```

## Project Structure
```
brandsdigger/
├── cmd/
│   └── main.go           # Main entry point of the application
├── internal/
│   ├── api/              # Handlers for REST endpoints
│   ├── services/         # Business logic and integrations
│   └── config/           # Configuration and environment handling
├── pkg/
│   ├── openai/           # OpenAI integration
│   └── godaddy/          # GoDaddy API client
├── go.mod                # Go module dependencies
├── go.sum                # Go module checksums
└── README.md             # Project documentation
```

## Dependencies
- **Go** (latest version recommended)
- **OpenAI API** for name generation
- **GoDaddy API** for domain verification

## License
This project is open-source and available under the MIT License.

## Author
Developed by [Umar Farooq Waheed]
