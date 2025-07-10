import librosa
import boto3
import psycopg2
import os
import argparse
import numpy as np

a = argparse.ArgumentParser(
    prog="track_upload"
)
a.add_argument("trackId")
args = a.parse_args()


# 54457b09-b0f7-4858-a84b-d75e75fa36b7
def database():
    user = os.environ["WORKER_PG_USER"]
    password = os.environ["WORKER_PG_PASSWORD"]
    host = os.environ["WORKER_PG_HOST"]
    port = os.environ["WORKER_PG_PORT"]
    db = os.environ['WORKER_PG_DATABASE']
    dn = f"postgres://{user}:{password}@{host}:{port}/{db}"
    return psycopg2.connect(dn)

def s3():
    return boto3.client("s3", endpoint_url=os.environ["INTERNAL_S3_ENDPOINT"])