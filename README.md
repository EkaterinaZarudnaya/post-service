# Post service

For this application, you will need Golang version 1.22.3 and Docker.

## Getting Started <a name = "getting_started"></a>

### Installing

#### Golang

Follow the [link](https://go.dev/doc/install) to download Golang version 1.22.3.

After installing Golang, add the GoPATH to your environment variables.

To install the dependencies follow the [link](https://go.dev/doc/modules/managing-dependencies) to get the details.

#### Docker
Follow the [link](https://docs.docker.com/desktop/install/windows-install/) to download and install Docker.

Check your docker-compose version.
```
docker-compose --version
```
### Running

### ! place .env file to the root project folder.

Use instructions in Makefile to run the app.
```
make run
```
#### The app will run locally using AWS RDS Postgres DB.

To run app in Docker containers, execute the following command in the root directory of your project:
```
docker-compose up -d
```
The app will run locally Postgres DB as container.
