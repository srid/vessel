# vessel

A tidy vessel for your [spews][https://github.com/heroku/spew].

## Getting started

For local development install Elm and [gb](https://github.com/constabulary/gb) followed by running `make`.

For deploying to Heroku:

```
heroku create
heroku buildpacks:set https://github.com/ddollar/heroku-buildpack-multi.git
heroku config:set GOVERSION=1.5
git push heroku master
```

## TODO

- [X] Deploy Go server to Heroku
- [X] End to end Go-Elm hello world
  - [X] With multibuildpack
- [ ] Stream plain logs to Elm
  - [ ] Broadcast channel and websocket abstraction
- [ ] Per app
- [ ] Analyze in real-time
