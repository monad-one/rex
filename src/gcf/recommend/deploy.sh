#!/bin/bash

gcloud functions deploy recommend --env-vars-file .env.yaml --runtime nodejs6 --trigger-http