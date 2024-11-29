-- roles table is
CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL
);

-- faculties table is
CREATE TABLE IF NOT EXISTS faculties (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    code VARCHAR(50) UNIQUE NOT NULL
);

-- classes table is
CREATE TABLE IF NOT EXISTS groups (
    id SERIAL PRIMARY KEY,
    faculty_id INT REFERENCES faculties (id) NOT NULL,
    class_code VARCHAR(50) UNIQUE NOT NULL
);

-- users table is
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    role_id INT REFERENCES roles (id),
    group_id INT REFERENCES groups (id),
    name VARCHAR(50) NOT NULL,
    surname VARCHAR(100) NOT NULL,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- create type for payment_status
CREATE TYPE payment_status_enum AS ENUM ('In Progress', 'Rejected', 'Accepted');

-- create type for payment_type
CREATE TYPE payment_type_enum AS ENUM ('I', 'II', 'I, II');

-- create type for course_year
CREATE TYPE course_year_enum AS ENUM ('1', '2', '3', '4', '5');

-- payments table is
CREATE TABLE IF NOT EXISTS payments (
    id SERIAL PRIMARY KEY,
    student_id INT REFERENCES users (id) ON DELETE CASCADE NOT NULL,
    payment_type payment_type_enum NOT NULL,
    payment_status payment_status_enum NOT NULL DEFAULT 'In Progress',
    course_year course_year_enum NOT NULL DEFAULT '1',
    check_photo VARCHAR(100) NOT NULL,
    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);