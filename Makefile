BINARY=engine
test: 
	go test -v -cover -covermode=atomic ./...

engine:
	go build -o ${BINARY} *.go


unittest:
	go test -short  ./...

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -t search-movie-service .

run:
	docker-compose up --build -d

stop:
	docker-compose down
