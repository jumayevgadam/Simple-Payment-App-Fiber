-- roles table is
CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    role VARCHAR(50) UNIQUE NOT NULL
);

-- permissions table is
CREATE TABLE permissions (
    id SERIAL PRIMARY KEY,
    permission_type VARCHAR(50)
);

CREATE TABLE times (
    id SERIAL PRIMARY KEY,
    start_year INT NOT NULL CHECK (start_year > 0),
    end_year INT NOT NULL CHECK (end_year > start_year)
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

CREATE MATERIALIZED VIEW payment_details_mv AS
SELECT 
    p.id AS payment_id,
    p.student_id,
    p.payment_type,
    p.payment_status,
    p.payment_amount,
    p.check_photo,
    p.uploaded_at,
    p.updated_at,
    u.name AS student_name,
    u.surname AS student_surname,
    g.course_year,
    t.start_year,
    t.end_year
FROM 
    payments p
JOIN 
    users u ON p.student_id = u.id
JOIN 
    groups g ON u.group_id = g.id
LEFT JOIN 
    times t ON p.time_id = t.id
WITH DATA;


CREATE UNIQUE INDEX idx_payment_details_mv_payment_id ON payment_details_mv (payment_id);
CREATE INDEX idx_payment_details_mv_student_name ON payment_details_mv (student_name);

