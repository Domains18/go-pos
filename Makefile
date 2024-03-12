hydrate:
	@templ generate

app:
	@templ generate
	@go build main
	@./notiwa

server:
	@air

container:
	@docker compose up