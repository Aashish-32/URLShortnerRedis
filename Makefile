compose:
		docker-compose up -d
rebuild:
		docker-compose build

# docker-compose down --rmi all
# docker system prune




.PHONY: Compose,rebuild
