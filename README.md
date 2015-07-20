## panoptic

This will be a collection of modules for accessing, storing and retrieving video frames and streams. It is designed to make writing live streaming and recording applications simple and powerful.

## Development

* panoptic is written in Go and is based on gstreamer and glib

### Packages

* `sudo apt-get install golang libgstreamer1.0-dev postgresql-9.3 libpq-dev postgis`

### Source
* `go get github.com/revmischa/panoptic`

### Database Init

```
sudo -u postgres createuser $USER
sudo -u postgres createdb -O$USER panoptic
echo "CREATE EXTENSION postgis" | sudo -u postgres psql panoptic
echo "CREATE EXTENSION pgcrypto" | sudo -u postgres psql panoptic
psql panoptic < schema.sql
```

## Design
_(Work in progress)_
