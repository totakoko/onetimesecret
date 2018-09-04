# Development Guide

## Prerequisites

Install these software:

- [Go](https://golang.org/) (*tested with 1.10.3*)
- [Node.js](https://nodejs.org/) (*tested with 10.8.0*)
- [Yarn](https://yarnpkg.com) (*tested with 1.9.4*)
- [Dep](https://golang.github.io/dep/) (*tested with 0.5.0*)
- [Docker](https://www.docker.com/) (*tested with 18.06.0-ce*)
- [Docker Compose](https://docs.docker.com/compose/) (*tested with 1.19.0*)


## Development

- Install the Node dependencies

    ```sh
    yarn install
    ```

- Then start the server with autorestart on file changes (reload is still manual).

    ```sh
    yarn start
    ```

The application is available at http://localhost:5000/.


## Building for production

- Build the assets for production.

    ```sh
    yarn run build
    ```

- Then build the server.

    ```sh
    go build
    ```

- Now a binary called `onetimesecret` should exist and the frontend content should be located in the `.build` folder.


## Design

### Goals

- [Keep the software simple](https://en.wikipedia.org/wiki/KISS_principle)
- Ease the deployment and monitoring by following the [12-factor app principles](https://12factor.net/)


### Compatibility

- [Web Cryptography API](https://caniuse.com/#feat=cryptography) : Firefox 61, Chrome 49
- [Async functions](https://caniuse.com/#search=async) : Firefox 61, Chrome 63


### JavaScript

The forms use traditional HTTP POST requests so that the application can work without JavaScript on the client side.
JavaScript enables client-side encryption so it is recommended though not enforced.


### HTML

[Pug](https://pugjs.org/) is used to simplify the markup language and ease the developments of pages.


### Styles

[Stylus](http://stylus-lang.com/) is used as a CSS preprocessor.
Usage of the BEM methodology is recommended by leveraging the `&` keyword of Stylus.
The pages are currently very simple so each page is viewed as a BEM component.

Traditionnal HTML elements such as links and buttons usually have their own style because they are used in multiple pages.

### Logo icons

- Edit the vector images in *frontend/src* with [Inkscape](https://inkscape.org).
- Then, inside the project root, use the following commands to export the images to png.

```sh
inkscape -z frontend/src/logo-icon.svg --export-png frontend/public/images/icon-512.png -w 512 -h 512
inkscape -z frontend/src/logo-icon.svg --export-png frontend/public/images/icon-192.png -w 192 -h 192
inkscape -z frontend/src/logo-icon.svg --export-png frontend/public/images/icon-32.png -w 32 -h 32
inkscape -z frontend/src/icon-lock.svg --export-plain-svg frontend/public/images/icon-lock.svg
```


## Project Structure

- .build: temporary folder containing the compiled assets
- common: contains interfaces that are used by other packages
- conf: contains the configuration of the application and its default values
- frontend: contains all frontend-related stuff
  - public: contains all assets that directly served to the browser
  - templates: contains the *pug* pages that are rendered in the browser
  - src: contains images in the source format (svg)
  - styles: contains the *stylus* files for styling the pages
- helpers: contains some helper functions that have no external dependencies and can be used by any package
- httpserver: contains the HTTP layer that serves the static content, render the pages and expose an API to interact with the store
- node_modules: contains the Node dependencies, managed with [yarn](https://yarnpkg.com/)
- store: contains the business and database layer responsible of persisting and retrieving the data from redis
- tests: contains End-to-End tests
- vendor: contains the go dependencies fetched by [dep](https://golang.github.io/dep/)
