
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
	godep save

test_api:
	#Create domain 1
	curl -is -X POST http://localhost:8080/api/v1/domain/foo
	@printf "\n"
	#Create domain 2
	curl -is -X POST http://localhost:8080/api/v1/domain/bar
	@printf "\n"
	#Get domain
	curl -is -X GET http://localhost:8080/api/v1/domains
	@printf "\n"
	#Create host 1
	curl -is -X POST http://localhost:8080/api/v1/host/bar?domain=foo
	@printf "\n"
	#Create host 2
	curl -is -X POST http://localhost:8080/api/v1/host/barry?domain=foo
	@printf "\n"
	#Get domain
	curl -is -X GET http://localhost:8080/api/v1/hosts?domain=foo
	@printf "\n"
	#Create var 1
	curl -is -X POST http://localhost:8080/api/v1/var/base/value/belongtous?domain=foo
	@printf "\n"
	#Create var 2
	curl -is -X POST http://localhost:8080/api/v1/var/allyour/value/base?domain=foo
	@printf "\n"
	#Get vars
	curl -is -X GET http://localhost:8080/api/v1/vars?domain=foo
	@printf "\n"
	#Create host var 1
	curl -is -X POST http://localhost:8080/api/v1/host/bar/var/feet/value/many?domain=foo
	@printf "\n"
	#Create host var 2
	curl -is -X POST http://localhost:8080/api/v1/host/bar/var/socks/value/black?domain=foo
	@printf "\n"
	#Get host vars
	curl -is -X GET http://localhost:8080/api/v1/host/bar/vars?domain=foo
	@printf "\n"
	#Create group
	curl -is -X POST http://localhost:8080/api/v1/group/hips?domain=foo
	@printf "\n"
	#Get groups
	curl -is -X GET http://localhost:8080/api/v1/groups?domain=foo
	@printf "\n"
	#Add group var
	curl -is -X POST http://localhost:8080/api/v1/group/hips/var/pants/value/jeans?domain=foo
	@printf "\n"
	#Add group var 2
	curl -is -X POST http://localhost:8080/api/v1/group/hips/var/boot/value/docs?domain=foo
	@printf "\n"
	#Get group vars
	curl -is -X GET http://localhost:8080/api/v1/group/hips/vars?domain=foo
	@printf "\n"
	#Add group host
	curl -is -X POST http://localhost:8080/api/v1/group/hips/host/stevo?domain=foo
	@printf "\n"
	#Add group host 2
	curl -is -X POST http://localhost:8080/api/v1/group/hips/host/johnnyk?domain=foo
	@printf "\n"
	#Get group hosts
	curl -is -X GET http://localhost:8080/api/v1/group/hips/hosts?domain=foo
	@printf "\n"
	#Add group host var 1
	curl -is -X POST http://localhost:8080/api/v1/group/hips/host/stevo/var/attack/value/crack?domain=foo
	@printf "\n"
	#Add group host var 2
	curl -is -X POST http://localhost:8080/api/v1/group/hips/host/stevo/var/shoe/value/nike?domain=foo
	@printf "\n"
	#Get group host vars
	curl -is -X GET http://localhost:8080/api/v1/group/hips/host/stevo/vars?domain=foo
	@printf "\n"
	#Get all
	curl -is -X GET http://localhost:8080/api/v1/all?domain=foo
	@printf "\n"
	#Get ansible groups
	curl -is -X GET http://localhost:8080/api/v1/ansible/groups?domain=foo
	@printf "\n"
