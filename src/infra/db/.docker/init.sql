CREATE USER user_test WITH PASSWORD 'user_test';
CREATE USER user_migration WITH PASSWORD 'user_migration';
GRANT ALL PRIVILEGES ON DATABASE "fx-fiber-db" to user_test;
GRANT ALL PRIVILEGES ON DATABASE "fx-fiber-db" to user_migration;
