from alpine:3.10

copy entrypoint.sh /entrypoint.sh

entrypoint ["/entrypoint.sh"]
