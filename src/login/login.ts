function checkCredentials(form: HTMLFormElement) {
    // @ts-ignore
    let data = new URLSearchParams(new FormData(form).entries())

    let res: boolean = checkContent(data)

    if (res) {
        fetch("/userLogin", {
            body: data,
            method: "POST"
        })
            .then(response => {
                if (response.ok) {

                    window.location.assign("/profile")

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
    } else {
        let errorField = document.getElementById("errorMessage")
        errorField.innerText = "Bitte Logindaten eingeben"
    }
}


function checkContent(data: URLSearchParams): boolean {

    let name: string = data.get('username')
    let pw: string = data.get('password')

    if (name.length == 0 || pw.length == 0)
        return false
    else
        return true

}