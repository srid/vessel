all:
	gb build
	elm make src/client/Vessel.elm --output static/index.html
