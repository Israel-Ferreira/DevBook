const criarUsuario = async (e) => {
    e.preventDefault()
    console.log("Dentro da função js")

    let nome = document.getElementById("name")
    let email = document.getElementById("email")
    let nick = document.getElementById("nick")

    let senha = document.getElementById("password")
    let confirmacaoSenha = document.getElementById("confirm-password")


    if (senha.value !== confirmacaoSenha.value) {
        alert("As senhas não coincidem")
        return
    }

    const request =  await fetch("/usuarios", {
        method: "POST",
        body: {
            nome: nome.value,
            email: email.value,
            nick: nick.value,
            senha: senha.value
        }
    })


    const resp =  await request

    if(resp.status == 201){
        console.log("Teste")
    }else{
        console.log("Erro ao submeter o formulario")
    }

}


$("#formulario-cadastro").on("submit", criarUsuario)