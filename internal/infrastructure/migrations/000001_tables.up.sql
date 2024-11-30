-- roles table is
CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL
);

-- permissions table is
CREATE TABLE permissions (
    id SERIAL PRIMARY KEY,
    permission_type VARCHAR(50)
);

-- role_permissions table is
CREATE TABLE role_permissions (
    role_id INT REFERENCES roles (id),
    permission_id INT REFERENCES permissions (id),
    PRIMARY KEY (role_id, permission_id)
);

-- faculties table is
CREATE TABLE IF NOT EXISTS faculties (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    faculty_code VARCHAR(50) UNIQUE NOT NULL
);

-- classes table is
CREATE TABLE IF NOT EXISTS groups (
    id SERIAL PRIMARY KEY,
    faculty_id INT REFERENCES faculties (id) NOT NULL,
    group_code VARCHAR(50) UNIQUE NOT NULL
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
CREATE TYPE payment_type_enum AS ENUM ('1', '2', '3');

-- create type for course_year
CREATE TYPE course_year_enum AS ENUM (1, 2, 3, 4, 5);

-- create type for semester
CREATE TYPE semester_enum AS ENUM ('first', 'second');

-- payments table is
CREATE TABLE IF NOT EXISTS payments (
    id SERIAL PRIMARY KEY,
    student_id INT REFERENCES users (id) ON DELETE CASCADE NOT NULL,
    payment_type payment_type_enum NOT NULL,
    payment_status payment_status_enum NOT NULL DEFAULT 'In Progress',
    course_year course_year_enum NOT NULL DEFAULT 1,
    check_photo VARCHAR(100) NOT NULL,
    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- indexes for payment table
CREATE INDEX idx_payments_status ON payments (payment_status);
CREATE INDEX idx_payments_student ON payments (student_id);
CREATE INDEX idx_payments_semester ON payments (semester);

-- indexes for users table
CREATE INDEX idx_group_id ON users (group_id) WHERE group_id IS NOT NULL;
CREATE INDEX idx_username ON users (username);
CREATE INDEX idx_name ON users (name);
CREATE INDEX idx_name_lower ON users (LOWER(name));
CREATE INDEX idx_surname ON users (surname);
CREATE INDEX idx_surname_lower ON users (LOWER(surname));