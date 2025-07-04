import argparse
import requests

parser = argparse.ArgumentParser(
                    prog='ProgramName',
                    description='What the program does',
                    epilog='Text at the bottom of help')

parser.add_argument("-b", "--host")
parser.add_argument("-m", "--mode")
parser.add_argument("-u", "--username")
parser.add_argument("-p", "--password")
args = parser.parse_args()
print(args.username, args.password)

if args.mode == "create":
    req = {
        "username": args.username,
        "password": args.password,
    }
    res = requests.post(args.host + "/api/signup", json=req)
    print(res)
    
elif args.mode == "login":
    req = {
        "username": args.username,
        "password": args.password,
    }
    res = requests.post(args.host + "/api/login", json=req)
    print(res)
    