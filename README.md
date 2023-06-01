## log into container
docker exec -it godockerDB /bin/sh
mysql -h localhost -u gorm -p gorm --default-character-set=cp850

## container sql backup and restore (https://gist.github.com/spalladino/6d981f7b33f6e0afe6bb)
#Backup
docker exec CONTAINER /usr/bin/mysqldump -u root --password=root DATABASE > backup.sql
# Restore
cat backup.sql | docker exec -i CONTAINER /usr/bin/mysql -u root --password=root DATABASE


WARNING: Image for service webserver-baseapp was built because it did not already exist. To rebuild this image you must use `docker-compose build` or `docker-compose up --build`.








