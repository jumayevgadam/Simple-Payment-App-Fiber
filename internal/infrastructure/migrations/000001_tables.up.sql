-- roles table is
CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    role VARCHAR(50) UNIQUE NOT NULL
);

INSERT INTO roles (
    role
) VALUES ('superadmin'), ('admin'), ('student');

CREATE TABLE times (
    id SERIAL PRIMARY KEY,
    start_year INT NOT NULL CHECK (start_year > 0),
    end_year INT NOT NULL CHECK (end_year > start_year)
);

-- faculties table is
CREATE TABLE IF NOT EXISTS faculties (
    id SERIAL PRIMARY KEY,
    faculty_name VARCHAR(50) UNIQUE NOT NULL,
    faculty_code VARCHAR(50) UNIQUE NOT NULL
);

-- classes table is
CREATE TABLE IF NOT EXISTS groups (
    id SERIAL PRIMARY KEY,
    faculty_id INT REFERENCES faculties (id) NOT NULL,
    group_code VARCHAR(50) NOT NULL,
    course_year INT NOT NULL CHECK (course_year > 0)
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

-- payments table is
CREATE TABLE IF NOT EXISTS payments (
    id SERIAL PRIMARY KEY,
    student_id INT REFERENCES users (id) ON DELETE CASCADE NOT NULL,
    time_id INT REFERENCES times (id),
    payment_type payment_type_enum NOT NULL,
    payment_status payment_status_enum NOT NULL DEFAULT 'In Progress',
    payment_amount INT NOT NULL,
    check_photo TEXT NOT NULL,
    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- indexes for payment table
CREATE INDEX idx_payments_status ON payments (payment_status);
CREATE INDEX idx_payments_student ON payments (student_id);
CREATE INDEX idx_payments_type ON payments (payment_type);

-- indexes for users table
CREATE INDEX idx_group_id ON users (group_id) WHERE group_id IS NOT NULL;
CREATE INDEX idx_username ON users (username);
CREATE INDEX idx_name ON users (name);
CREATE INDEX idx_name_lower ON users (LOWER(name));
CREATE INDEX idx_surname ON users (surname);
CREATE INDEX idx_surname_lower ON users (LOWER(surname));

-- indexes for faculties table 
CREATE INDEX idx_faculty_name ON faculties (faculty_name);
CREATE INDEX idx_faculty_code ON faculties (faculty_code);

-- indexes for groups table 
CREATE INDEX idx_group_code ON groups (group_code);

INSERT INTO users (
    role_id, name, surname, username, password
) VALUES (
    1, 'Gadam', 'Jumayev', 'hypergadam', '$2a$10$O1GSNPZgig1p.qLQj2Ibve8k1aDBey9Ll9I4D/XFxObPaUev3Fvhm'
) RETURNING id;

ALTER TABLE times
ADD COLUMN is_active BOOLEAN DEFAULT FALSE;

ALTER TABLE payments
ADD CONSTRAINT unique_student_time_payment UNIQUE (student_id, time_id, payment_type);

ALTER TABLE faculties 
ADD COLUMN faculty_index INT;

ALTER TABLE groups 
ADD COLUMN group_index INT;