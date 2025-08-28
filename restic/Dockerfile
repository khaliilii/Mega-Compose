FROM alpine:latest

RUN apk add --no-cache restic bash openssh

WORKDIR /backup
ENTRYPOINT ["restic"]


#docker build -t my-restic .

#docker run --rm -v /data:/data -v /backup:/backup my-restic backup /data

#run restic:
#docker compose run --rm restic
#backup list:
#docker compose run --rm restic snapshots
#restore:
#docker compose run --rm restic restore latest --target /restore
#schdule restic:
#0 2 * * * docker compose run --rm restic


#docker run --rm \
#  -v /data:/data \
#  -v /backup:/backup \
#  restic/restic \
#  backup /data
  