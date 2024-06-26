dev:
	make -j docker-up 

dev-stop:
	docker compose down

docker-up:
	docker compose up -d

api-setup:
	cd api && go mod download

api-start:
	cd api && air

setup:
	make -j api-setup web-setup

web-setup:
	cd web && npm install

web-start:
	cd web && npm run dev

web-start-h:
	cd web && npm run dev:host
