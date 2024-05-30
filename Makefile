
templ:
	templ generate -watch -proxy=http://localhost:8090

tailwind:
	tailwindcss --config configs/tailwind.config.js \
		-i configs/input.css \
		-o assets/css/styles.css \
		-w \
		--minify

