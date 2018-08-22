# One Time Secret

*Licensed under [AGPL-3](#license)*

A clone of the very useful [One-Time Secret](https://onetimesecret.com) but with a simpler design ([demo](onetimesecret.dreau.fr)).

Features:
- share a secret that can be read only one time
- (TODO) secure with client-side encryption: the server has no idea what the secrets are
- minimalist and beautiful design: no unnecessary features, no advertisement for third-party services or the product itself, no freemium
- mobile-friendly

Apart from its usefulness for sharing secrets with other people, this project is mainly used to experiment best development practices such as: automated tests, continuous deployment, configuration management, monitoring...


## Usage

The simplest option is to use the [docker-compose](https://docs.docker.com/compose/) file below to get a service running:

- First create a folder named `ots` and copy the content below into a file called `docker-compose.yml`.

```yaml
version: '3.3'
services:
  onetimesecret:
    image: registry.dreau.fr/home/onetimesecret:latest
    restart: unless-stopped
    ports:
      - "8080:5000"
    environment:
      OTS_STORE_ADDR: redis:6379
      OTS_PUBLICURL: http://localhost:8080/
    networks:
      - onetimesecret
  redis:
    image: redis:4.0-alpine
    restart: unless-stopped
    ports:
      - "6379:6379"
    networks:
      - onetimesecret

networks:
  onetimesecret:
```

- Inside the folder, run `docker-compose up -d`.
- After a few seconds or minutes, depending of your internet connection, you will be able to access OTS locally at http://localhost:8080/.


## Development

See the [development guide](./DEVELOPMENT.md).


## License

The code is licensed under the [GNU Affero General Public License version 3](./LICENSE.md).
You can find a human-readable summary on [tldrlegal.com](https://tldrlegal.com/license/gnu-affero-general-public-license-v3-(agpl-3.0)).
