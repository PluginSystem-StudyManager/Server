function checkRegistration(form: HTMLFormElement) {

    let colorError = '#DC143C'
    let errorField = document.getElementById("errorMessage")
    // @ts-ignore
    let data = new URLSearchParams(new FormData(form).entries())

    let password1 = data.get('Password')
    let passwort2 = data.get('PasswordAgain')

    let result: boolean = comparePassword(password1, passwort2);

    if (result == false) {
        console.log('Das War wohl nix')

        let pwField = document.getElementsByName('Password')
        pwField[0].style.backgroundColor = colorError
        let pwAgainField = document.getElementsByName('PasswordAgain')
        pwAgainField[0].style.backgroundColor = colorError

        errorField.innerText = "passwords do not match"

    } else {

        fetch("/checkUserName", {
            body: data,
            method: "POST"
        })

    }


}

function comparePassword(pwd: string, pwd2: string): boolean {

    let nummer: number = pwd.localeCompare(pwd2)
    if (nummer == 0) {
        return true
    } else return false;
}