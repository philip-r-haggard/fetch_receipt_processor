# Fetch Receipt Processor README

## Getting Started

### Clone the Repository

```bash
git clone https://github.com/philip-r-haggard/fetch_receipt_processor.git
cd fetch_receipt_processor
```

### Install Dependencies

To install needed dependencies use these commands in a bash terminal:

```bash
go mod init fetch_receipt_processor
go get github.com/google/uuid
go mod tidy
```

### Run the application

To run the application use this command in a bash terminal:

```bash
go run .
```

### Test the application

I recommend using Postman to test the application.

For the Process Receipts endpoint, add a new request in Postman. Set the method to POST and use this url:

```bash
http://localhost:8080/receipts/process
```

In the body of the request, check the raw option and add your JSON similar to this:

```bash
{
    "retailer": "Target",
    "purchaseDate": "2022-01-02",
    "purchaseTime": "13:13",
    "total": "1.25",
    "items": [
        {"shortDescription": "Pepsi - 12-oz", "price": "1.25"}
    ]
}
```

For the Get Points endpoint, add a new request in Postman. Set the method to GET and use this url:

```bash
http://localhost:8080/receipts/{id}/points
```

where {id} is the id received from the Process Receipts endpoint.

### Unit test the application

To run the unit tests for the application use this command in a bash terminal:

```bash
go test -v
```