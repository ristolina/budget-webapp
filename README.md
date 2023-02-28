# budget-webapp
Budget WebApp is my final project in the Cloud Native Computing Course (https://ya.se/yrkeshogskola/cloud-native-computing/).

# How to try it for yourself?
1. Clone the repo or download the files from https://github.com/ristolina/budget-webapp.
2. Create the SECRETS folder and files with the different secret keys as referenced in the docker-compose.yaml
3. Run docker compose build
4. Run docker compose up
5. Connect to docker mariaDB instance and create the 'expenses' table and insert test data according to DB/init-db.sql
6. Webapp is available on http://localhost:5000. Backend is available on http://localhost:5001.