let cart = [];

function loadSuppliers() {
  apiGet("/suppliers").done(res => {
    res.forEach(s => {
      $("#supplier").append(`<option value="${s.id}">${s.name}</option>`);
    });
  });
}

function loadItems() {
  apiGet("/items").done(res => {
    res.forEach(i => {
      $("#item").append(`<option value="${i.id}">${i.name} (stok: ${i.stock})</option>`);
    });
  });
}

function loadPurchases() {
  apiGet("/purchasings").done(res => {
    $("#purchaseTable").empty();
    res.forEach(p => {
      $("#purchaseTable").append(`
        <tr>
          <td>${p.id}</td>
          <td>${p.supplier?.name || "-"}</td>
          <td>${p.user?.username || "-"}</td>
          <td>${p.grand_total}</td>
          <td>
            <button class="btn btn-danger btn-sm delete" data-id="${p.id}">Hapus</button>
          </td>
        </tr>
      `);
    });
  });
}


$("#addItem").click(() => {
  const itemID = $("#item").val();
  const itemText = $("#item option:selected").text();
  const qty = parseInt($("#qty").val());

  if (!itemID || qty <= 0) {
    Swal.fire("Error", "Item dan qty wajib diisi", "error");
    return;
  }

  cart.push({ item_id: parseInt(itemID), qty });

  $("#cartTable").append(`
    <tr>
      <td>${itemText}</td>
      <td>${qty}</td>
      <td>
        <button class="btn btn-sm btn-danger remove">Hapus</button>
      </td>
    </tr>
  `);

  $("#qty").val("");
});

$("#cartTable").on("click", ".remove", function () {
  const idx = $(this).closest("tr").index();
  cart.splice(idx, 1);
  $(this).closest("tr").remove();
});


$("#submitOrder").click(() => {
  const supplierID = $("#supplier").val();

  if (!supplierID || cart.length === 0) {
    Swal.fire("Error", "Supplier dan item wajib diisi", "error");
    return;
  }

  apiPost("/purchasings", {
    supplier_id: parseInt(supplierID),
    items: cart
  })
  .done(() => {
    Swal.fire("Success", "Purchase berhasil", "success");
    cart = [];
    $("#cartTable").empty();
    loadPurchases();
  })
  .fail(xhr => {
    Swal.fire(
      "Error",
      xhr.responseJSON?.message || "Gagal membuat purchase",
      "error"
    );
  });
});


$("#purchaseTable").on("click", ".delete", function () {
  const id = $(this).data("id");

  Swal.fire({
    title: "Yakin?",
    text: "Purchase akan dihapus dan stok dikembalikan",
    icon: "warning",
    showCancelButton: true
  }).then(result => {
    if (result.isConfirmed) {
      apiDelete(`/purchasings/${id}`)
        .done(() => {
          Swal.fire("Deleted", "Purchase dihapus", "success");
          loadPurchases();
        });
    }
  });
});


$(document).ready(() => {
  if (!localStorage.getItem("token")) {
    window.location.href = "login.html";
  }

  loadSuppliers();
  loadItems();
  loadPurchases();
});
