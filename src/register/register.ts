const colorError = '#DC143C'
const errorField = document.getElementById("errorMessage")

function checkRegistration(form: HTMLFormElement) {

    // @ts-ignore
    let data = new URLSearchParams(new FormData(form).entries())

    let password1 = data.get('Password')
    let passwort2 = data.get('PasswordAgain')

    let noEmptyFields: boolean = allFieldsFilledOut()

    let result: boolean = comparePassword(password1, passwort2);

    if (result == false) {

        let pwField = document.getElementsByName('Password')
        pwField[0].style.backgroundColor = colorError

        let pwAgainField = document.getElementsByName('PasswordAgain')
        pwAgainField[0].style.backgroundColor = colorError

        showErrorMessage("passwords do not match")
    }
    if (noEmptyFields) {

        fetch("/checkUserName", {
            body: data,
            method: "POST"
        }).then(function (response) {

            if (response.ok) {
                return response.json()
            } else {

            }
        }).then(function (json) {

            if (json.Fehlercode == 5) {
                showErrorMessage(json.Fehlermeldung)
            } else if (json.Fehlercode == 0) {  // kein Fehler

                window.location.assign("/")
            }
        })

    }
}


function comparePassword(pwd: string, pwd2: string): boolean {

    let nummer: number = pwd.localeCompare(pwd2)
    if (nummer == 0) {
        return true
    } else return false;
}


function allFieldsFilledOut(): boolean {

    let fieldNames = ['FirstName', 'LastName', 'UserName', 'EMail', 'Password', 'PasswordAgain'];
    let emptyFields = 0

    for (let i = 0; i < fieldNames.length; i++) {

        emptyFields = fieldIsEmpty(fieldNames[i], emptyFields)
    }
    if (emptyFields == 0)
        return true
    else
        return false
}


function fieldIsEmpty(value: string, counter: number): number {

    let textField = <HTMLInputElement>document.getElementsByName(value)[0]
    if (textField.value.length == 0) {

        textField.style.backgroundColor = colorError
        showErrorMessage("Es wurden nicht alle Felder ausgefüllt")
        counter++
    } else {
        textField.style.backgroundColor = ''
    }

    if (counter == 0) {
        showErrorMessage("")
    }
    return counter
}


function showErrorMessage(message: string) {

    errorField.innerText = message
}

