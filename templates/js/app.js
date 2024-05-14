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
})