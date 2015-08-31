# vessel

A tidy vessel for your spews.

## Work in progress

```
go get github.com/constabulary/gb/...
make
```

Pushing to heroku:

```
heroku create
heroku buildpacks:set https://github.com/ddollar/heroku-buildpack-multi.git
heroku buildpacks:set https://github.com/paxan/heroku-buildpack-gb.git
heroku config:set GOVERSION=1.5
git push heroku master
```

## TODO

- [X] Deploy Go server to Heroku
- [ ] End to end Go-Elm hello world
  - [ ] With multibuildpack
- [ ] Get plain logs to Elm
- [ ] Per app
- [ ] Analyze in real-time
