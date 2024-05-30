# FitWave

Fitwawe is a tool that lets you fetch your Strava data, 
create custom dashboards using SQL, and view your workouts 
with tables and charts. It helps you understand your training 
better and make informed decisions about your fitness.

## Features
- **Data Integration:** Connect with your Strava account and get your workout data.
- **Custom Dashboards:** Use SQL to create dashboards that show the information you want.
- **Visualization Tools:** See your workout data in tables and charts.
- **User-Friendly Interface:** Easy to use and navigate.
- **Real-Time Updates:** Keep your dashboards updated with the latest data from Strava.

## Install the Project

There are a few options to install the project:

- [Download a release binary](#release-binaries)
- [Download a docker image](#docker-image)
- [Build a binary from source](#building-from-source)

### Release binaries

You can download the release binary for your system from the [releases page](https://github.com/bahattincinic/fitwave/releases).

### Docker image

To pull the Docker image:

```bash
# Pull the latest image
docker pull bahattincinic/fitwave

# Or specify the image by tag
docker pull bahattincinic/fitwave[:tag]
```

### Building from source

**1. Clone the Repository**

```
git clone git@github.com:bahattincinic/fitwave.git
cd fitwave
```

**2. Build Backend**

```
make
```

**2. Build Frontend**

```
cd ui
npm install
```

#### Running the Application

**1. Run Backend**

Run project with the following command;

    ./fitwave

If you need to override the configuration, create a `.env` file.

**2. Run Frontend (`different terminal tab`)**

    cd ui
    npm run serve

## Environment Variables

| Field Name            | Type    | Default    | Options/Examples                                                                           |
|-----------------------|---------|------------|--------------------------------------------------------------------------------------------|
| ENV                   | string  | local      | - local<br/>- production                                                                   |
| LOG_LEVEL             | string  | debug      | - debug<br/>- info<br/>- warn<br/>- error<br/>- panic<br/>- fatal                          |
| LOG_OUTPUT            | string  | stdout     | - stdout<br/>- /foo/bar/fitwave.log                                                        |
| DATABASE_DSN          | string  | fitWave.db | - fitWave.db<br/>- host=localhost user=postgres password=postgres dbname=fitwave port=5432 |
| DATABASE_TYPE         | string  | sqlite     | - mysql<br/>- postgresql<br/>- sqlite                                                      |
| DATABASE_AUTO_MIGRATE | boolean | true       | - true<br/>- false                                                                         |
| API_PORT              | integer | 9000       |                                                                                            |

## Production Build

Build Frontend

```
make build-frontend
```

Build Backend

```
make GCFLAGS="-tags=prod"
```


## Swagger API Documentation

To view the API Swagger documentation, please visit the following link:

[http://localhost:9000/api/docs/](http://localhost:9000/api/docs/)

Updating Swagger

    make swaggen
