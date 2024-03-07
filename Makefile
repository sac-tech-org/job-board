dev:
	make -j docker-up api-start web-start

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
