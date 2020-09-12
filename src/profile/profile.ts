function AccountSettins() {
    let x = document.getElementById("Plugin");
    let y = document.getElementById("Account");
    x.style.display = "none"
    y.style.display = "block"
    let btnPlugins = document.getElementById("btnPlugins")
    let btnAccount = document.getElementById("btnAccount")
    btnPlugins.classList.remove("selected")
    btnAccount.classList.add("selected")
    window.location.hash = "#account"
}

function Plugins() {
    let x = document.getElementById("Account");
    let y = document.getElementById("Plugin");
    x.style.display = "none"
    y.style.display = "flex"
    let btnPlugins = document.getElementById("btnPlugins")
    let btnAccount = document.getElementById("btnAccount")
    btnPlugins.classList.add("selected")
    btnAccount.classList.remove("selected")
    window.location.hash = "#plugins"
}

function copyPermanentToken() {
    let item = <HTMLInputElement>document.getElementById("permanent_token")
    item.select()
    document.execCommand("copy")
    // only https: navigator.clipboard.writeText(token).then(r => console.log("Copied: " + navigator.clipboard.readText()))
}

document.addEventListener('DOMContentLoaded', function () {
    if (window.location.hash.substr(1) === "plugins") {
        Plugins()
    } else {
        AccountSettins()
    }
}, false);
