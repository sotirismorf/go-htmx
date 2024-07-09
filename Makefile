css-dev:
	@npm run dev

css-build:
	@npm run build

image:
	@docker build --tag go-htmx:latest .
