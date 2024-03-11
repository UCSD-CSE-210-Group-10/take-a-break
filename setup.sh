sudo docker stop my-postgres-db react-frontend go-backend
sudo docker rm my-postgres-db go-backend react-frontend
sudo docker image rm -f postgres-image go-backend-image react-frontend-image
sudo docker network rm servernet

sudo docker network create --subnet=172.18.0.0/16 servernet

cd sql
sudo docker build -t postgres-image .
sudo docker run --net servernet --ip 172.18.0.2 -d -v /var/run/postgresql:/var/run/postgresql --name my-postgres-db postgres-image

sleep 10

cd ../backend
sudo docker build -t go-backend-image .
sudo docker run --net servernet -p 8080:8080 --ip 172.18.0.3 -d --name go-backend go-backend-image

sleep 10

cd ../frontend
sudo docker build -t react-frontend-image .
sudo docker run --net servernet -p 3000:3000 --ip 172.18.0.4 -d --name react-frontend react-frontend-image
