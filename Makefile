dev:
	@make -j templ tailwind

templ:
	@templ generate --watch --proxy="http://localhost:8080" --cmd="go run ."

tailwind:
	@npx tailwindcss -i ./input.css -o ./assets/styles.css --watch
