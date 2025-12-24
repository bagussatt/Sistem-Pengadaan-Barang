$("#loginForm").on("submit", function (e) {
  e.preventDefault();

  apiPost("/login", {
    username: $("#username").val(),
    password: $("#password").val(),
  })
  .done(function (res) {
    localStorage.setItem("token", res.token);

    Swal.fire("Success", "Login berhasil", "success")
      .then(() => window.location.href = "purchase.html");
  })
  .fail(function (xhr) {
    Swal.fire(
      "Error",
      xhr.responseJSON?.error || "Login gagal",
      "error"
    );
  });
});
