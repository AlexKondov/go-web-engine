#!/bin/bash

clear
rm basaigbook
go build --tags="mgo"
./basaigbook -driver mongo -datasource "127.0.0.1"