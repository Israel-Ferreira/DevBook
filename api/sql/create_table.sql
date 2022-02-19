CREATE DATABASE IF NOT EXISTS DevBook;

DROP TABLE IF EXISTS usuarios;
DROP TABLE IF EXISTS seguidores;

CREATE TABLE usuarios (
    id INT AUTO_INCREMENT PRIMARY KEY, 
    nome VARCHAR(55) NOT NULL,
    nick VARCHAR(55) NOT NULL UNIQUE,
    email VARCHAR(55) NOT NULL UNIQUE,
    senha VARCHAR(512) NOT NULL,
    criadoEm TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE = INNODB;



CREATE TABLE seguidores (
    usuario_id INT NOT NULL,
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id)
    ON DELETE CASCADE,

    seguidor_id INT NOT NULL,
    FOREIGN KEY (seguidor_id) REFERENCES usuarios(id)
    ON DELETE CASCADE,

    PRIMARY KEY(usuario_id, seguidor_id)
);