all: server client
	@true

server:
	gb build

client:
	elm make src/client/Vessel.elm --output static/index.html
