hydrate:
	@templ generate

app:
	@templ generate
	@go build main
	@./notiwa

serve:
	@templ generate
	@go run main.go

container:
	@docker compose up