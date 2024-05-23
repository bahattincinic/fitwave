# FitWave

Virtualize Your Strava Workouts

## Setup

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

## Running the Application

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
| ENV                   | string  | local      | - local<br/>- testing<br/>- production                                                     |
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
