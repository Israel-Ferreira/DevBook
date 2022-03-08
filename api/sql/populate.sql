insert into usuarios (nome, nick, email, senha)
values 
(
    "Luiz Pareto",
    "luiz.pareto",
    "luiz.pareto@telerj.io",
    "$2a$10$0R3NxunwS0eJmqsIzbNnsOpYl35uukV2uWRo1R5EJnDDgZI7dumkq"
),
(
    "Usuario 1",
    "usuario.1",
    "usuario1@example.com",
    "$2a$10$0R3NxunwS0eJmqsIzbNnsOpYl35uukV2uWRo1R5EJnDDgZI7dumkq"
),
(
    "Cachorro Caramelo",
    "meloww.dog",
    "meloww.dog@example.com",
    "$2a$10$0R3NxunwS0eJmqsIzbNnsOpYl35uukV2uWRo1R5EJnDDgZI7dumkq"
);

insert into seguidores (usuario_id, seguidor_id)
values 
(1, 3),
(1, 4),
(1, 5);


INSERT INTO publicacoes (titulo, conteudo, autor_id)
VALUES
("Publicação do Usuário 2", "Minha primeira publicação Uhuu", 2),
("Publicação do Usuário 1", "Minha primeira publicação Uhuu", 1),
("Publicação do Usuário 3", "Minha primeira publicação Uhuu", 3);