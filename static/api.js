const API_KEY = "1312o4ajfkaj";
const USER_NAME = "c@b.com";
const PASSWORD = "pw";

const postOptions = data => {
  return {
    method: "POST", // or 'PUT'
    body: JSON.stringify(data), // data can be `string` or {object}!
    headers: {
      "Content-Type": "application/json"
    }
  };
};

const postApi = (url, payload) => {
  return fetch(url, postOptions(payload)).then(response => {
    console.log("API - POST", response);
    if (response.status == 200) {
      return response.json();
    } else {
      return {
        success: false,
        error: response.status,
        message: response.statusText
      };
    }
  });
};

const getOptions = {
  method: "GET", // or 'PUT'
  body: null,
  headers: {
    "Content-Type": "application/json"
  }
};

const getApi = url => {
  return fetch(url, getOptions).then(response => {
    console.log("API - GET", response);
    if (response.status == 200) {
      return response.json();
    } else {
      return {
        success: false,
        error: response.status,
        message: response.statusText
      };
    }
  });
};

const display = id => {
  const app = document.getElementById("app-page");
  const loading = document.getElementById("loading");
  const login = document.getElementById("login");
  switch (id) {
    case "app":
      loading.style.display = "none";
      login.style.display = "none";
      app.style.display = "block";
      break;
    case "login":
      app.style.display = "none";
      loading.style.display = "none";
      login.style.display = "block";
      window.document.getElementById("username").value = "";
      window.document.getElementById("password").value = "";
      break;
    case "loading":
      login.style.display = "none";
      app.style.display = "none";
      loading.style.display = "block";
      break;
  }
};

const auth = () => {
  getApi("/auth")
    .then(data => {
      console.log(data);
      if (data.success) {
        localStorage.setItem("data", JSON.stringify(data));
        display("app");
        renderPage();
      } else {
        switch (data.error) {
          case 401:
            display("login");
            break;
          default:
            display("login");
            break;
        }
      }
    })
    .catch(error => {
      console.log("RESPONSE ERROR (rare) " + error);
    });
};

const login = (username, password) => {
  showStatusMessage("Logging In.....");
  postApi("/login", {
    username,
    password
  })
    .then(data => {
      console.log(data);
      if (data.success) {
        localStorage.setItem("data", JSON.stringify(data));
        hideStatusMessage();
        display("app");
        renderPage();
      } else {
        switch (data.error) {
          case 401:
            showStatusMessageToast("Invalid Username/password", 5000);
            break;
          default:
            break;
        }
      }
    })
    .catch(error => {
      console.log("RESPONSE ERROR (rare) " + error);
      return false;
    });
};

const logout = id => {
  showStatusMessage("Logging Out.....");
  postApi("/logout", { id })
    .then(data => {
      console.log(data);
      localStorage.removeItem("data");
      localStorage.removeItem("purchases");
      hideStatusMessage();
      display("login");
    })
    .catch(error => {
      console.log("RESPONSE ERROR (rare) " + error);
      hideStatusMessage();
      display("login");
    });
};

function showStatusMessage(msg) {
  window.document.getElementById("status-message").innerHTML = msg;
  window.document.getElementById("status").style.display = "block";
}
function hideStatusMessage() {
  window.document.getElementById("status").style.display = "none";
}

function showStatusMessageToast(msg, delay) {
  showStatusMessage(msg);
  window.setTimeout(function() {
    hideStatusMessage();
  }, delay);
}
