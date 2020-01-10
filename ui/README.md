## Front-end widget

Created using React and was bootstrapped with [Create React App](https://github.com/facebook/create-react-app).

### Running widget
Using [Makefile](https://en.wikipedia.org/wiki/Makefile)

####To run in development
```bash
make run-dev
```

Will run development server on ```http://localhost:3000```

Available config file [config.js](public/config.js)

####Docker setup
To build docker image and start container:

```bash
make start
```

Additionally it's possible only build a docker image without running it:

```bash
make build
```

Docker container also will run on ```http://localhost:3000``` same port as development
environment

 