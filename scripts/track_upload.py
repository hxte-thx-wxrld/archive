import tempfile

conn = database()
cur = conn.cursor()
cur.execute("SELECT uri, trackname, release_date, artistid, status FROM uploads WHERE id=%s", (args.trackId,))
data = cur.fetchone()
#conn.close()
c = s3()

print(args.trackId)

cur.execute("UPDATE public.uploads SET status='analyzing' WHERE id::text = %s", (args.trackId,))
conn.commit()

with tempfile.NamedTemporaryFile() as tfile:
    #c.download_file()    
    print("Fetched Data:", data)
    c.download_fileobj('inbox', "track/" + args.trackId, tfile)
    
    
    y, sr = librosa.load(tfile.name)

    features = {
        "tempo": librosa.feature.tempo(y=y, sr=sr)[0],
        "duration": librosa.get_duration(y=y, sr=sr),
        "zcr": np.mean(librosa.feature.zero_crossing_rate(y)),
        "rms": np.mean(librosa.feature.rms(y=y)),
        "centroid": np.mean(librosa.feature.spectral_centroid(y=y, sr=sr)),
        "rolloff": np.mean(librosa.feature.spectral_rolloff(y=y, sr=sr)),
        "flatness": np.mean(librosa.feature.spectral_flatness(y=y)),
        "mfcc": np.mean(librosa.feature.mfcc(y=y, sr=sr, n_mfcc=13), axis=1).tolist(),
    }
    
    print(features)
    
    c.upload_fileobj(tfile, "tracks", args.trackId + ".wav")
    cur.execute("INSERT INTO music (id, title, artist_id, public_url, release_date) VALUES(%s, %s, %s, %s, %s)", (args.trackId, data[1], data[3], "/tracks/" + args.trackId + ".wav", data[2]))
    print("Inserted into music library")
    
    cur.execute("INSERT INTO public.analysis (trackid, tempo, zcr, rms, centroid, rolloff, flatness, duration) VALUES(%s, %s, %s, %s, %s, %s, %s, %s)", (args.trackId, features["tempo"].item(), features["zcr"].item(), features["rms"].item(), features["centroid"].item(), features["rolloff"].item(), features["flatness"].item(), features["duration"]))
    print("Inserted Analysis Data")
    
    cur.execute("UPDATE public.uploads SET status='finished' WHERE id::text = %s", (args.trackId,))
    conn.commit()
    print("Update Status")
    

print("worker finished")