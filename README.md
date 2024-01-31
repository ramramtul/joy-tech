# joy-tech

Preparation

install gin to run the application
go install github.com/codegangsta/gin@latest

install mux
go get github.com/gorilla/mux v1.8.1

install testify
go get github.com/stretchr/testify v1.8.4

install
sudo apt-get install build-essential


Running Application

to run the application you can use
gin -i --appPort 8080 --port 3000 run main.go

your good to go! just hit using curl or postman


Testing Application

to test all system
go test ./... -coverprofile=coverage.out

to view code coverage
go tool cover -html=coverage.out


Application Limitation:
can only pick a book that have been searched (BookList)
    - so you have to search using GET /book/list
    - then submit a book that shown on the list
    - Edition Number is openlibrary_work