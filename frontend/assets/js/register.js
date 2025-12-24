$("#registerForm").on("submit", function (e) {
  e.preventDefault();

  const data = {
    username: $("#username").val(),
    password: $("#password").val(),
    role: $("#role").val() 
  };

  apiPost("/register", data)
    .done(function () {
      Swal.fire("Success", "Register berhasil", "success")
        .then(() => window.location.href = "login.html");
    })
    .fail(function (xhr) {
      Swal.fire(
        "Error",
        xhr.responseJSON?.error || "Register gagal",
        "error"
      );
    });
});
