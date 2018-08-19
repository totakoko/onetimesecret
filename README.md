# One Time Secret

Clone of the very useful [One-Time Secret](https://onetimesecret.com) but with a simpler design.

This project is mainly used to experiment best development practices, automated tests, continuous deployment...


## Prerequisites

- go (*tested with 1.10.3*)
- node (*tested with 10.8.0*)
- yarn (*tested with 1.9.4*)

## Development

Start the process for compiling Pug to HTML and Stylus to CSS:

    ```sh
    yarn run gulp dev
    ```

Start the server:

    ```sh
    go run main.go
    ```


### JavaScript

The application requires no JavaScript on the client side.
The forms use traditional HTTP POST requests and almost no JavaScript is currently running.


### HTML

Pug is used to simplify the markup language and ease the developments of pages.


### Styles

Usage of the BEM methodology is recommended through Stylus powerful syntax.
The pages are currently very simple so each page is viewed as a BEM component.


TO BE UPDATED

Recompile CSS when a change is made to the stylus:

    ```sh
    npx pug -w templates/*.pug
    npx stylus -w public/assets/main.styl -o public/assets/main.css

    docker run -it --rm --user $UID:$GID \
        -v $PWD/public:/inputfiles \
        -v $PWD/public:/outputfiles \
        gruen/stylus /inputfiles/assets/main.styl -o /outputfiles/assets/main.css
    ```

## Project Structure

- public: contains all assets that map to
- templates: contains the *pug* pages that are rendered in the browser
- styles: contains the *stylus* files for styling the pages
- store: contains the business and database layer responsible of persisting and retrieving the data from redis


## Logo icons

- Edit the vector images in *frontend/src* with Inkscape.
- Then use the following commands to export the images to png.

    ```
    inkscape -z -e frontend/public/images/icon-512.png -w 512 -h 512 frontend/src/logo-icon.svg
    inkscape -z -e frontend/public/images/icon-192.png -w 192 -h 192 frontend/src/logo-icon.svg
    inkscape -z -e frontend/public/images/icon-32.png -w 32 -h 32 frontend/src/logo-icon.svg
    ```
