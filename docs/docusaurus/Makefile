PROJECT_PATH := $(shell pwd)

static-website:
	rm -rf build
	yarn run build
	rm -rf ../{img,docs,*.js,*.xml,*.LICENSE,*.html,*.css}
	cp -R build/* ..
