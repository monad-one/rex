#!/bin/bash

gcloud dataproc jobs submit pyspark gs://danicat/algs/als.py --cluster rex-dataproc-cluster --region us-east1 -- $MYSQL_HOST $MYSQL_DB $MYSQL_USER $MYSQL_PASSWORD