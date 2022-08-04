#!/bin/bash

docker run --rm -v PWD:/srv/app -w /srv/app golang:alpine go $@