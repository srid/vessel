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
heroku buildpacks:set https://github.com/paxan/heroku-buildpack-gb.git
heroku config:set GOVERSION=1.5
git push heroku master
```

## TODO

- [ ] End to end Go-Elm hello world
- [ ] Get plain logs to Elm
- [ ] Per app
- [ ] Analyze in real-time
