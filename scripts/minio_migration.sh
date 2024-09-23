#!/bin/bash

mc config host add minio http://minio:9000 $MINIO_ROOT_USER $MINIO_ROOT_PASSWORD
mc mb --ignore-existing minio/"$VIDEO_BUCKET"
mc anonymous set public minio/"$VIDEO_BUCKET"
mc mb --ignore-existing minio/"$IMAGE_BUCKET"
mc anonymous set public minio/"$IMAGE_BUCKET"


for file in data/*.mp4; do
    if [ -f "$file" ]; then
        mc put "$file" minio/"$VIDEO_BUCKET"
    fi
done
