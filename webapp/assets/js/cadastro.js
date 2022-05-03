$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(evento) {
    evento.preventDefault();

    if ($('#password').val() != $('#confirm-password').val()) {

        return;
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: $('#name').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: $('#password').val()
        }
    }).done(function () {
        alert("Usuário Criado com sucesso")
    }).fail(function () {
        alert("Erro ao criar o usuário")
    });
}