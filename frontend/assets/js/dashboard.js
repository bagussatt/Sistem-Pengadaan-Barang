if (!localStorage.getItem("token")) {
  window.location.href = "login.html";
}

$("#itemForm").on("submit", function (e) {
  e.preventDefault();

  apiRequest({
    url: "/items",
    method: "POST",
    data: {
      name: $("#itemName").val(),
      stock: Number($("#itemStock").val()),
      price: Number($("#itemPrice").val())
    }
  })
  .done(() => {
    Swal.fire("Berhasil", "Item berhasil ditambahkan", "success");
    this.reset();
  })
  .fail(err => {
    const msg =
      err?.responseJSON?.message ||
      err?.responseJSON?.error ||
      "Gagal menambahkan item";

    if (err.status === 401) {
      Swal.fire("Session habis", "Silakan login ulang", "warning")
        .then(() => {
          localStorage.removeItem("token");
          window.location.href = "login.html";
        });
      return;
    }

    Swal.fire("Error", msg, "error");
  });
});
