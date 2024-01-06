package main

//go:generate -command bundler go run ./cmd/bundler
//go:generate templ generate
//go:generate bundler
//go:generate ./tailwindcss -i input.css -o static/main.css --minify
