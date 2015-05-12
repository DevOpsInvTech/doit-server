
test:
	-rm _test_tmp/*.db
	godep go test -v .

build:
	godep go build .

test_api:
	#Create domain
	curl -v -X POST http://localhost:8080/api/v1/domain/foo
	#Create host
	curl -v -X POST http://localhost:8080/api/v1/host/bar?domain=foo
	#Create host var 1
	curl -v -X POST http://localhost:8080/api/v1/host/bar/var/feet/value/many?domain=foo
	#Create host var 2
	curl -v -X POST http://localhost:8080/api/v1/host/bar/var/socks/value/black?domain=foo
	#Get host vars
	curl -v -X GET http://localhost:8080/api/v1/host/bar/vars?domain=foo
