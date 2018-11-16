#!/bin/bash

FUNCTION=`gcloud functions describe recommend | grep url | cut -d' ' -f4`
echo Calling $FUNCTION

curl -X POST $FUNCTION -H "Content-Type:application/json" -d '{ "user_id": 11, "movie_id": 10 }' | jq