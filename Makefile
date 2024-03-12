hydrate:
	@templ generate

app:
	@templ generate
	@go build main
	@./notiwa

serve:
	@templ generate --watch
	@air

container:
	@docker compose up