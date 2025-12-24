$("#registerForm").on("submit", function (e) {
  e.preventDefault();

  apiRequest("POST", "/api/register", {
    username: $("#username").val(),
    password: $("#password").val(),
    role: $("#role").val()
  })
  .done(() => {
    alert("Register berhasil, silakan login");
    window.location.href = "login.html";
  })
  .fail(err => showError(err));
});
