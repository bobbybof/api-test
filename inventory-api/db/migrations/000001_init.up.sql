CREATE TABLE users (
    id SERIAL NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    password TEXT NOT NULL,
    email TEXT,
    created_at TIMESTAMP default(now()),
    updated_at TIMESTAMP default(now())
);

CREATE TABLE products (
    id SERIAL NOT NULL PRIMARY KEY,
    price double precision default(0::double precision) NOT NULL,
    name TEXT not null, 
    description TEXT,
    created_at TIMESTAMP default(now()),
    updated_at TIMESTAMP default(now())
);

CREATE TABLE clients (
    id SERIAL NOT NULL PRIMARY KEY,
    service_id INT NOT NULL,
    user_id INT NOT NULL,
    collaboration_start_date TIMESTAMP NOT NULL,
    contract_value double precision default(0::double precision) NOT NULL,
    created_at TIMESTAMP default(now()),
    updated_at TIMESTAMP default(now()),

    FOREIGN KEY (user_id) REFERENCES users(id)   
);