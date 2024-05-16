document.addEventListener("DOMContentLoaded", function () {
    const btn1 = document.getElementById("btn1")
    const lng1 = document.getElementById("lng1")

    btn1.addEventListener("click", function () {
        lng1.classList.toggle("language-active")
    })
    document.addEventListener("click", function () {
        if (!btn1.contains(event.target) && event.target !== lng1) {
            lng1.classList.remove("language-active")
        }
    })

    const btn2 = document.getElementById("btn2")
    const lng2 = document.getElementById("lng2")

    btn2.addEventListener("click", function () {
        lng2.classList.toggle("language-active")
    })
    document.addEventListener("click", function () {
        if (!btn2.contains(event.target) && event.target !== lng2) {
            lng2.classList.remove("language-active")
        }
    })

    const button1 = document.getElementById("button1")
    const language1 = document.getElementById("language1")

    button1.addEventListener("click", function () {
        language1.classList.toggle("language-active")
    })
    document.addEventListener("click", function () {
        if (!button1.contains(event.target) && event.target !== language1) {
            language1.classList.remove("language-active")
        }
    })

    const button2 = document.getElementById("button2")
    const language2 = document.getElementById("language2")

    button2.addEventListener("click", function () {
        language2.classList.toggle("language-active")
    })
    document.addEventListener("click", function () {
        if (!button2.contains(event.target) && event.target !== language2) {
            language2.classList.remove("language-active")
        }
    })

    const account = document.getElementById("account")
    const close = document.getElementById("close")
    const auth = document.getElementById("auth")
    const over = document.getElementById("over")


    account.addEventListener("click", function () {
        auth.classList.toggle("auth-window-active")
    })
    if (!account.contains(event.target) && event.target !== auth) {
        auth.classList.remove("auth-window-active")
    }
    close.addEventListener("click", function () {
        auth.classList.remove("auth-window-active")
    })

    account.addEventListener("click", function () {
        over.classList.toggle("over-active")
    })
    if (!account.contains(event.target) && event.target !== auth) {
        auth.classList.remove("auth-window-active")
    }
    close.addEventListener("click", function () {
        over.classList.remove("over-active")
    })


})