CREATE DATABASE IF NOT EXISTS DevBook;


DROP TABLE IF EXISTS publicacoes;

DROP TABLE IF EXISTS seguidores;

DROP TABLE IF EXISTS usuarios;


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

CREATE TABLE publicacoes (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    titulo VARCHAR(50) NOT NULL,
    conteudo VARCHAR(300) NOT NULL,
    autor_id int not null,
    FOREIGN KEY (autor_id) REFERENCES usuarios(id)
    ON DELETE CASCADE,

    curtidas INT NOT NULL DEFAULT 0,
    criadaEm TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);