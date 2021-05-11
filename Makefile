.PHONY: clean run-users run-passwords docker-build-users docker-build-passwords docker-run-users docker-run-passwords docker-create-network docker-run-users-internal docker-run-passwords-internal
.EXPORT_ALL_VARIABLES:

MAINFILES = main.go
SHELL=/bin/bash
SERVICE_PASSWORD_ADDRESS=http://localhost:8001

# 1.Let's run the services locally
run-users:
	@echo "===>  Running Users locally"
	@cd ./service-users && go run main.go

run-passwords:
	@echo "===>  Running Passwords locally"
	@cd ./service-passwords && go run main.go

# 2. Let's build container images
docker-build-users:
	@echo "===>  Building Users Docker Image"
	@cd ./service-users && 	CGO_ENABLED=0 GOOS=linux GARCH=amd64 go build -o app-linux $(MAINFILES) && docker build -t service-users .

docker-build-passwords:
	@echo "===>  Building Passwords Docker Image"
	@cd ./service-passwords && 	CGO_ENABLED=0 GOOS=linux GARCH=amd64 go build -o app-linux $(MAINFILES) && docker build -t service-passwords .


# 3. Let's run the services inside Docker (Opps, but we have a network problem)
docker-run-users:
	@echo "===>  Running Users Docker Image"
	@docker run -ti --rm --name=service-users -p 8000:8000 service-users:latest

docker-run-passwords:
	@echo "===>  Running Passwords Docker Image"
	@docker run -ti --rm --name=service-passwords -p 8001:8001 service-passwords:latest

# 4. Ok, let's create an internal network and try again
docker-create-network:
	@echo "===>  Creating Internal network"
	@docker network create --driver bridge internal

docker-run-users-internal:
	@echo "===>  Running Users Docker Image"
	@docker run -ti --rm --name=service-users --network internal -p 8000:8000 service-users:latest

docker-run-passwords-internal:
	@echo "===>  Running Passwords Docker Image"
	@docker run -ti --rm --name=service-passwords --network internal -p 8001:8001 service-passwords:latest

# 5. Further notes
# - What if we can have 2 password services?
# - What if we can have more?
# - Enterprise system designs

# Clean up 
clean:
	@echo "===> Cleaning"
	@docker system prune -a
