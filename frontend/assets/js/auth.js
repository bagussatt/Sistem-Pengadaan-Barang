$("#loginForm").on("submit", function (e) {
  e.preventDefault();

  apiRequest("POST", "/api/login", {
    username: $("#username").val(),
    password: $("#password").val()
  })
  .done(res => {
    localStorage.setItem("token", res.token);
    window.location.href = "dashboard.html";
  })
  .fail(err => showError(err));
});
