CREATE TABLE users (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username TEXT NOT NULL UNIQUE,
    company TEXT NOT NULL,
    phone_number TEXT NOT NULL
);