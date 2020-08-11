
function checkRegistration(form: HTMLFormElement) {

    // @ts-ignore
    let data = new URLSearchParams(new FormData(form).entries())
    console.log("ich war mal hier")

    fetch("/checkUserName", {
        body: data,
        method: "POST"
    })





}