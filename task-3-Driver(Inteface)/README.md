# Task-3-Driver(Inteface)
Extract data from postgres database and parse into json object

# Requirements(Python)
- python
- postgres

# Setup database
Check 'setup.sql' file to setup test db

# Setup and run app(Python)
> To be run on a windows os
- run `python -m venv venv`
- run `./venv/Script/activate`
- run `pip install -r requirements.txt`
- replace placeholder values in 'sample.env' with original values
- rename 'sample.env' to '.env'
- run 'app.py'

# Requirements(GO)
- GO
- postgres

# Setup and run app(GO)
> To be run on a windows os

- Install Go
- run `go env -w GO111MODULE=auto`
- run `go mod init app.go`
- run `go mod tidy`
- run `go mod init goenv`
- run `go get github.com/joho/godotenv`
- run `go run .\app.go`
