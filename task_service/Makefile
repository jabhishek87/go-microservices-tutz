# phony commands
.PHONY: help doc

help: 			### Show this help.
	${INFO} "Prints all the options"
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

doc:			### generate api docs
	${INFO} "Generating api documentation"
	@swag init -g server.go



# Cosmetics
YELLOW := "\e[1;33m"
NC := "\e[0m"

# Shell Functions
INFO := @bash -c '\
  printf $(YELLOW); \
  echo "=> $$1"; \
  printf $(NC)' SOME_VALUE