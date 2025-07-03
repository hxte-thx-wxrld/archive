#!/bin/sh
pg_dump-17 --clean -h 10.0.0.4 -U rillo -s --no-privileges --no-owner --if-exists -f db/test.sql