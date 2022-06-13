


const LoginUsuario = (e) => {
    e.preventDefault()



    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email: $('#email').val(),
            password: $('#password').val()
        }
    }).done(function (resp) {
        alert("Usuário Logado com sucesso sucesso")
        localStorage.setItem("token", resp["token"])
        window.location = "/home"
    }).fail(function (err) {
        console.log(err)
        alert("Erro ao fazer  o login do usuário")
    });
}

$("#login").on('submit', LoginUsuario)