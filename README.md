# One Time Secret

Clone of the very usefule [One Time Secret]() but without



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
