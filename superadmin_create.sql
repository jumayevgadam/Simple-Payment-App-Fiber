INSERT INTO users (
    role_id, name, surname, username, password
) VALUES (
    1, 'Gadam', 'Jumayev', 'hypergadam', '$2a$10$O1GSNPZgig1p.qLQj2Ibve8k1aDBey9Ll9I4D/XFxObPaUev3Fvhm'
) RETURNING id;