-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
BEGIN;

CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    last_logged_in_at TIMESTAMP,
    user_type VARCHAR(50)
    );

CREATE TABLE properties (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    room_count INTEGER NOT NULL
);

CREATE TABLE rooms(
    id SERIAL PRIMARY KEY,
    room_type VARCHAR(255) NOT NULL,
    bed_type VARCHAR(255) NOT NULL,
    rent INTEGER NOT NULL,
    property_id INTEGER REFERENCES properties(id) ON DELETE CASCADE 
);

CREATE TABLE bookings(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    room_id INTEGER REFERENCES rooms(id) ON DELETE CASCADE,
    booked_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
    booked_for DATE
);


COMMIT;
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
BEGIN;


DROP TABLE IF EXISTS bookings;

DROP TABLE IF EXISTS rooms;

DROP TABLE IF EXISTS properties;

DROP TABLE IF EXISTS users;

COMMIT;

