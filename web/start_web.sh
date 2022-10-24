#!/bin/bash

>> micro .env

PORT=3000

>> ctrl + s
>> ctrl + q

cd app

go build

./start_app
