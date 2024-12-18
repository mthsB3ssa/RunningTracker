# Subir os contêineres com build
up:
	@echo "Iniciando os contêineres com build..."
	docker-compose up --build

# Parar e remover os contêineres
down:
	@echo "Parando os contêineres..."
	docker-compose down

# Reiniciar os contêineres
restart:
	@echo "Reiniciando os contêineres com build..."
	docker-compose down
	docker-compose up --build
