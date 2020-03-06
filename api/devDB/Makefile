TAG='e-portfolio-db'
NAME='e-portfolio-db-container'
build:
	docker build -t $(TAG) .

run:
	docker run -d --rm -e POSTGRES_PASSWORD=pass --name $(NAME) -p 5432:5432 $(TAG)

psql:
	psql -h 127.0.0.1 -U postgres

stop:
	docker stop $(NAME)
