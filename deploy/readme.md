# deploy

- make p
- update config on docker-compose.yml
- move `deploy` folder to your server
- docker compose up -d

## Update Config docker-compose.yml

- Change `traefik.example.dev` to your domain
- `your@gmail.com` to your email
- `backend` to your project name
- `/ex-prefix` to your prefix path

- You can use image or build from dockerfile

    ```yaml
    version: "3.1"

    services:

    ...

    backend:
        # build:
        #   context: .
        #   dockerfile: "../Dockerfile"
        image: bb

    ...

    ```
