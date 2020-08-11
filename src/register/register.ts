function checkRegistration(form: HTMLFormElement) {

    // @ts-ignore
    let data = new URLSearchParams(new FormData(form).entries())

    let password1 = data.get('Password')
    let passwort2 = data.get('PasswordAgain')

    let result: boolean = comparePassword(password1, passwort2);

    if (result == false) {
        console.log('Das War wohl nix')
        return
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