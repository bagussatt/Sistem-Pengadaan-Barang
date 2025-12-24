if (!localStorage.getItem("token")) {
  window.location.href = "login.html";
}

// Handler untuk tambah Item
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
    // Reload item table
    apiRequest({ url: "/items", method: "GET" }).done(items => {
      let html = `
        <tr>
          <th>ID</th><th>Name</th><th>Stock</th><th>Price</th>
        </tr>`;
      items.forEach(i => {
        html += `
          <tr>
            <td>${i.ID}</td>
            <td>${i.Name}</td>
            <td>${i.Stock}</td>
            <td>${i.Price}</td>
          </tr>`;
      });
      $("#itemTable").html(html);
    });
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

// Handler untuk tambah Supplier
$("#supplierForm").on("submit", function (e) {
  e.preventDefault();

  apiRequest({
    url: "/suppliers",
    method: "POST",
    data: {
      name: $("#supplierName").val(),
      email: $("#supplierEmail").val(),
      address: $("#supplierAddress").val()
    }
  })
  .done(() => {
    Swal.fire("Berhasil", "Supplier berhasil ditambahkan", "success");
    this.reset();
    // Reload supplier table
    apiRequest({ url: "/suppliers", method: "GET" }).done(data => {
      let html = `
        <tr>
          <th>ID</th><th>Name</th><th>Email</th><th>Address</th>
        </tr>`;
      data.forEach(s => {
        html += `
          <tr>
            <td>${s.ID}</td>
            <td>${s.Name}</td>
            <td>${s.Email || "-"}</td>
            <td>${s.Address || "-"}</td>
          </tr>`;
      });
      $("#supplierTable").html(html);
    });
  })
  .fail(err => {
    const msg =
      err?.responseJSON?.message ||
      err?.responseJSON?.error ||
      "Gagal menambahkan supplier";

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

// Handler untuk logout
$("#logoutBtn").on("click", function () {
  Swal.fire({
    title: "Logout",
    text: "Apakah Anda yakin ingin logout?",
    icon: "question",
    showCancelButton: true,
    confirmButtonText: "Ya",
    cancelButtonText: "Batal"
  }).then((result) => {
    if (result.isConfirmed) {
      localStorage.removeItem("token");
      window.location.href = "login.html";
    }
  });
});
