function checkUsername(e) {
    const value = e.target.value;
    if(!validUsername(value)) {
        // Display error message
    }
}

function checkPassword(e) {
    const value = e.target.value;
    if(!validPassword(value)) {
        // Display error message
    }
}

function validUsername(username) {
    return false;
}

function validPassword(password) {
    return false;
}

function tryLogin() {
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

    if(!(validUsername(username) && validPassword(password))) {
        // Display error message
        return;
    }
}