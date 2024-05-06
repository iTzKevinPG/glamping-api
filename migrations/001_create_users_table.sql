-- Crear la tabla de usuarios
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    address TEXT,
    phone VARCHAR(20),
    password VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

-- Agregar un índice en el campo username para acelerar las búsquedas
CREATE INDEX idx_username ON users (username);

-- Opcionalmente, agregar un índice para el campo email si se espera que las búsquedas por email sean comunes
CREATE INDEX idx_email ON users (email);

