
# Weather Stations

A simple Go module for managing weather stations and temperature measurements - porazdeljeni sistemi, naloga 5.

## Features

- Add weather stations.
- Add temperature measurements.
- Calculate minimum, maximum, and average temperatures.
- Display measurements and weather reports.
- Configure temperature limits and the minimum number of measurements using command-line flags.

## Installation


go get github.com/EulersIdentity33/vremenske-postaje@v1.0.0


## Usage

Run the application with default settings:

go run ./main

Run with custom settings:

go run ./main --minTemp=-10 --maxTemp=30 --stMeritev=4

## Package

The `vreme` package provides functions for managing weather stations.

- `DodajPostajo` – adds a new weather station.
- `DodajMeritev` – adds a temperature measurement.
- `IzpisMeritev` – displays all measurements.
- `IzpisPorocila` – displays weather reports.

## Documentation

[View the vreme package on pkg.go.dev](https://pkg.go.dev/github.com/EulersIdentity33/vremenske-postaje/vreme)
