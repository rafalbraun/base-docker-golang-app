## connect to host via ssh
`ssh -p 22 USERNAME@HOSTNAME`

## deploy (copy files) via scp
`sshpass -p PASSWORD scp -r ./* USERNAME@HOSTNAME:/home/username`
`docker-compose up`

## log into container
`docker exec -it CONTAINER /bin/sh`

## log into database locally
`mysql -h localhost -u gorm -p gorm --default-character-set=cp850`

## container sql backup [here](https://gist.github.com/spalladino/6d981f7b33f6e0afe6bb)
`docker exec CONTAINER /usr/bin/mysqldump -u root --password=root DATABASE > backup.sql`

## container sql restore [here](https://gist.github.com/spalladino/6d981f7b33f6e0afe6bb)
`cat backup.sql | docker exec -i CONTAINER /usr/bin/mysql -u root --password=root DATABASE`

## rebuild image
`docker-compose build` or `docker-compose up --build`

## launch locally
```
docker-compose up mysql-baseapp
make run
```
