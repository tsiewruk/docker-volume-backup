# docker-volume-backup

The purpose of the script is to create a snapshot of the Docker volume.

## Installation
`sudo curl -SL  -o /usr/local/bin/docker-volume-backup`
`sudo chmod +x /usr/local/bin/docker-volume-backup`

## Requirements
- `python` - minimum version 3.10

## Usage
`Create volume backup:
docker-volume-back create <volume_name> <destination_path/backup_name.tar>
Example:
docker-volume-backup create prometheus_volume /opt/backup/prometheus.tar

Restore volume backup
docker-volume-backup restoree <destination_path/backup_name.tar> <volume_name>
Example:
docker-volume-backup restore /opt/backup/prometheus.tar prometheus_volume`

## Cron Usage
`* * * * * /usr/local/bin/docker-volume-backup <volume_name> <destination_path/backup_name.tar>`
