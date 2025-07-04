from sql import connection
import boto3
import pathlib
import tempfile

s3 = boto3.resource('s3')

def get_music_in_db():
    with connection.cursor() as cur:
        cur.execute("select id, public_url from music m where public_url is not null")
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
            
def download_from_s3(fileobj, object_name):
    s3.Bucket("tracks").Object(object_name).download_fileobj(fileobj)
            
if __name__ == "__main__":
#        upload_to_s3(track)
    for track in get_music_in_db():
        with tempfile.NamedTemporaryFile() as f:
            p = pathlib.Path(track[1])
            print(p.name)
            download_from_s3(f, p.name)