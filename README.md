# joy-tech

## Preparation

install gin to run the application <br>
go install github.com/codegangsta/gin@latest

install mux <br>
go get github.com/gorilla/mux v1.8.1

install testify <br>
go get github.com/stretchr/testify v1.8.4

install <br>
sudo apt-get update <br>
sudo apt-get install build-essential


## Running Application

to run the application you can use <br>
gin -i --appPort 8080 --port 3000 run main.go

your good to go! just hit using curl or postman


## Testing Application

to test all system with coverage <br>
go test -short -coverprofile coverage.out -v ./...

to view code coverage in detail <br>
go tool cover -html=coverage.out


## Application Limitation:
can only pick a book that have been searched (BookList) <br>
    - so you have to search using GET /book/list <br>
    - then submit a book that shown on the list <br>
    - Edition Number is openlibrary_work <br>
