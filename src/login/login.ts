function checkCredentials(form: HTMLFormElement) {
    let data = new FormData(form)

    fetch("/userLogin", {
        body: data,
        method: "POST"
    })
        .then(response => {
            console.log(response)
            if (response.ok) {

                window.location.assign("/test")

            } else {
                let errorField = document.getElementById("errorMessage")

                if (response.status == 401) {  // Unauthorized User
                    errorField.innerText = "Username / Password is wrong"
                }
                if (response.status == 500) {  //Internal Server Error
                    errorField.innerText = "Server is currently not available"
                }
            }
        })

}