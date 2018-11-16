#!/bin/bash

gcloud beta functions deploy recommend --env-vars-file .env.yaml --runtime nodejs6 --trigger-http