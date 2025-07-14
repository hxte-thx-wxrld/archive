# archive
The backbone of the htw music catalog

## Installation and Setup
Requirements:
- Docker
- Minio Client

To get started you need to setuo a new serviceaccount, aka Access Keys on the minio instance. To do so, start the minio instance if it is not already running by using:
`docker compose up minio`

If you have used the compose file in the repository, your user account is `root/rootpassword`. Running the following commands sets up our minio environment and adds the needed account:
```bash
minio 
```