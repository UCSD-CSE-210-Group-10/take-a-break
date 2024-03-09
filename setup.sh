sudo docker network create --subnet=172.18.0.0/16 servernet

cd sql
sudo docker build -t postgres-image .
sudo docker run --net servernet --ip 172.18.0.2 -d -v /var/run/postgresql:/var/run/postgresql --name my-postgres-db postgres-image

sleep 10

cd ../backend
sudo docker build -t go-backend-image .
sudo docker run --net servernet --ip 172.18.0.3 -d --name go-backend go-backend-image

