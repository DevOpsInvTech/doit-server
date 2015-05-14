
test:
	-rm _test_tmp/*.db
	godep go test -v .

clean:
	-rm _test_tmp/*.db

build:
	godep go build .

all:
	godep go install -a
	godep go build .

run:
	./doit-server -c test_configs/test-config.yml -s

update_deps:
	godep update github.com/DevOpsInvTech/doittypes

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
