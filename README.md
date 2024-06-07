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

To run the unit tests for the application use this command in a bash terminal:

```bash
go test -v
```