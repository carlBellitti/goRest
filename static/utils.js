//Comment 
let users = [
  { user: 'barney', age: 36, active: true },
  { user: 'fred',  age: 40, active: false },
  { user: 'travis', age: 37, active: true}
];

result = _.filter(users, function(o) { return o.active; });
console.log(result);


let myPurchases = [];
let appData = null;
auth()
function loginToApp () {
    console.log("Login To App")
    username = window.document.getElementById("username").value
    password = window.document.getElementById("password").value
    console.log("USERNAME",username, "PASSWORD", password)
    login(username, password)
}


function logoutOfApp () {
    console.log("logout of App")
    logout()
}
function getMyPurchases () {
    console.log("getPurchases")
    showStatusMessage("Getting Purchases .....");
    getApi("/purchases?id=" + appData.id, {})
        .then(data => {
        console.log(data);
        if (data.success) {
            myPurchases = data.purchases;
            hideStatusMessage();
            renderPurchases();
        } else {
            switch (data.error) {
            case 401:
                display("login")
                break;
            default:
                display("login")
                break;
            }
        }
        })
        .catch(error => {
        console.log("RESPONSE ERROR (rare) " + error);
        return false;
        });

}
function renderPage() {
    data = JSON.parse(localStorage.getItem("data")).data
    appData = data;
    document.getElementById("id").value = data.id
    document.getElementById("full-name").value = data.fullname
    availableList = document.getElementById("available-movie-list")
    data.moviedata.map(movie => {
        const li = document.createElement("li");
        li.appendChild(document.createTextNode(movie));
        li.setAttribute("class", "list-group-item");
        availableList.appendChild(li);
    })
}
function renderPurchases() {
    myPurchaseList = document.getElementById("purchased-movie-list")
    myPurchases.map(purchase => {
        const li = document.createElement("li");
        li.appendChild(document.createTextNode(purchase));
        li.setAttribute("class", "list-group-item");
        myPurchaseList.appendChild(li);
    })
}