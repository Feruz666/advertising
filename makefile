bd:
	docker run --name advertisingDB -p 54322:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=secret -d postgres:latest