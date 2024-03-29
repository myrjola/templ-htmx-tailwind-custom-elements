# templ-htmx-tailwind-custom-elements
A Go example of Templ, HTMX, Tailwind and Custom Elements

## Quickstart

Install [Tailwind](https://tailwindcss.com/blog/standalone-cli) standlalone:

```
# Example for macOS x64
curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-x64
chmod +x tailwindcss-macos-x64
mv tailwindcss-macos-x64 tailwindcss
```

Install [Templ](https://templ.guide/quick-start/installation):

```
go install github.com/a-h/templ/cmd/templ@latest
```

Run code generation:

```
go generate ./...
```


## Development

Watch for changes with Tailwind to update styles:

```
./tailwindcss -i input.css -o static/main.css --watch
```

Start dev environment with hot reload:

```
templ generate --watch --cmd="./start_dev_server.sh" --proxy="http://127.0.0.1:8080"
```