from sql import connection
import boto3
import pathlib

s3 = boto3.resource('s3')

def get_music_in_db():
    with connection.cursor() as cur:
        cur.execute("select id, filepath from music m")
        return cur.fetchall()
    
def get_music_in_db_local_only():
    with connection.cursor() as cur:
        cur.execute("select id, filepath from music m where public_url is null")
        return cur.fetchall()
        
def upload_to_s3(track):
    with connection.cursor() as cur:
        with open(track[1], 'rb') as data:
            p = pathlib.Path(track[1])

            #print(filename, p.suffix, p)
            s3_name = track[0] + p.suffix
            s3.Bucket('tracks').put_object(Key=s3_name, Body=data)
            cur.execute("update music set public_url = %s where id = %s", ("/tracks/" + s3_name, track[0]))
            connection.commit()
            print(track[1], "=>", s3_name)
            #print("Finished ", track)
            
if __name__ == "__main__":
    for track in get_music_in_db_local_only():
        upload_to_s3(track)