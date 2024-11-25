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