let cart = [];
let itemsData = [];

if (!localStorage.getItem("token")) {
  window.location.href = "login.html";
}

apiRequest({ url: "/suppliers", method: "GET" }).done(data => {
  data.forEach(s => {
    $("#supplier").append('<option value="' + s.ID + '">' + s.Name + '</option>');
  });
});

apiRequest({ url: "/items", method: "GET" }).done(data => {
  itemsData = data;
  data.forEach(i => {
    var stockText = i.Stock === 0 ? "(STOK HABIS)" : "(Stock: " + i.Stock + ")";
    var disabled = i.Stock === 0 ? "disabled" : "";
    $("#item").append('<option value="' + i.ID + '" data-stock="' + i.Stock + '" data-price="' + i.Price + '" data-name="' + i.Name + '" ' + disabled + '>' + i.Name + ' ' + stockText + '</option>');
  });
});

$("#item").on("change", function() {
  var selectedOption = $(this).find(":selected");
  var stock = selectedOption.data("stock");
  var price = selectedOption.data("price");

  if (stock !== undefined) {
    if (stock === 0) {
      $("#itemInfo").html('<span class="text-danger fw-bold">STOK HABIS!</span>');
    } else {
      $("#itemInfo").text("Stock: " + stock + " | Price: Rp " + Number(price).toLocaleString());
    }
  } else {
    $("#itemInfo").text("");
  }
});

$("#addItem").click(function() {
  var itemId = $("#item").val();
  var qty = Number($("#qty").val());
  var selectedOption = $("#item").find(":selected");
  var itemName = selectedOption.data("name");
  var itemPrice = Number(selectedOption.data("price"));
  var itemStock = Number(selectedOption.data("stock"));

  if (!itemId || !qty) {
    Swal.fire({ icon: "warning", title: "Data Tidak Lengkap", text: "Silakan pilih item dan masukkan qty", confirmButtonColor: "#ffc107" });
    return;
  }

  if (itemStock === 0) {
    Swal.fire({ icon: "error", title: "Stok Habis", text: "Maaf, stok untuk item ini sudah habis.", confirmButtonColor: "#dc3545" });
    return;
  }

  if (qty > itemStock) {
    Swal.fire({ icon: "error", title: "Qty Terlalu Banyak", text: "Qty melebihi stock yang tersedia (" + itemStock + ")", confirmButtonColor: "#dc3545" });
    return;
  }

  var existingIndex = cart.findIndex(function(c) { return c.item_id === Number(itemId); });
  if (existingIndex !== -1) {
    var newQty = cart[existingIndex].qty + qty;
    if (newQty > itemStock) {
      Swal.fire({ icon: "error", title: "Qty Terlalu Banyak", text: "Total qty melebihi stock yang tersedia (" + itemStock + ")", confirmButtonColor: "#dc3545" });
      return;
    }
    cart[existingIndex].qty = newQty;
  } else {
    cart.push({
      item_id: Number(itemId),
      qty: qty,
      name: itemName,
      price: itemPrice
    });
  }

  renderCart();
  $("#qty").val("");
  $("#item").val("").trigger("change");
});

function renderCart() {
  var html = "";
  var total = 0;

  if (cart.length === 0) {
    html = "<tr><td colspan='5' class='text-center text-muted'>Cart kosong</td></tr>";
  } else {
    cart.forEach(function(c, i) {
      var subtotal = c.price * c.qty;
      total += subtotal;
      html += "<tr>" +
        "<td>" + c.name + "</td>" +
        "<td>Rp " + c.price.toLocaleString() + "</td>" +
        "<td>" + c.qty + "</td>" +
        "<td>Rp " + subtotal.toLocaleString() + "</td>" +
        "<td><button class='btn btn-danger btn-sm remove' data-i='" + i + "'>Hapus</button></td>" +
        "</tr>";
    });
  }
  $("#cartTable").html(html);
  $("#cartTotal").text("Rp " + total.toLocaleString());
}

$(document).on("click", ".remove", function () {
  var index = $(this).data("i");
  cart.splice(index, 1);
  renderCart();
});

$("#submitOrder").click(function() {
  var supplierId = $("#supplier").val();

  if (!supplierId) {
    Swal.fire({ icon: "warning", title: "Supplier Belum Dipilih", text: "Silakan pilih supplier terlebih dahulu", confirmButtonColor: "#ffc107" });
    return;
  }

  if (cart.length === 0) {
    Swal.fire({ icon: "warning", title: "Cart Kosong", text: "Silakan tambahkan item ke cart terlebih dahulu", confirmButtonColor: "#ffc107" });
    return;
  }

  var orderItems = cart.map(function(c) {
    return { item_id: c.item_id, qty: c.qty };
  });

  apiRequest({
    url: "/purchasings",
    method: "POST",
    data: {
      supplier_id: Number(supplierId),
      items: orderItems
    }
  })
  .done(function() {
    var total = $("#cartTotal").text();
    Swal.fire({
      icon: "success",
      title: "Purchase Berhasil!",
      html: "Purchase order berhasil dibuat.<br><br>Total: " + total + "<br>Stock akan diupdate otomatis.<br><br>Halaman akan di-refresh dalam 2 detik...",
      timer: 2000,
      showConfirmButton: false
    });
    cart = [];
    renderCart();
    loadPurchases();
    setTimeout(function() { location.reload(); }, 2000);
  })
  .fail(function(err) {
    Swal.fire({
      icon: "error",
      title: "Gagal Membuat Purchase",
      text: err.responseJSON?.message || "Terjadi kesalahan",
      confirmButtonColor: "#dc3545"
    });
  });
});

function loadPurchases() {
  apiRequest({ url: "/purchasings", method: "GET" }).done(function(data) {
    var html = "";
    if (data.length === 0) {
      html = "<tr><td colspan='5' class='text-center text-muted'>Belum ada purchase</td></tr>";
    } else {
      data.forEach(function(p) {
        var supplierName = p.Supplier ? p.Supplier.Name : "-";
        var userName = p.User ? p.User.Username : "-";
        var total = p.GrandTotal ? p.GrandTotal.toLocaleString() : 0;
        var date = p.Date ? new Date(p.Date).toLocaleDateString() : "-";
        html += "<tr>" +
          "<td>" + p.ID + "</td>" +
          "<td>" + supplierName + "</td>" +
          "<td>" + userName + "</td>" +
          "<td>Rp " + total + "</td>" +
          "<td><button class='btn btn-info btn-sm view-detail' data-id='" + p.ID + "'>Detail</button></td>" +
          "</tr>";
      });
    }
    $("#purchaseTable").html(html);
  });
}

$(document).on("click", ".view-detail", function () {
  var id = $(this).data("id");
  apiRequest({ url: "/purchasings/" + id, method: "GET" }).done(function(p) {
    var details = "";
    if (p.Details && p.Details.length > 0) {
      p.Details.forEach(function(d) {
        var itemData = itemsData.find(function(i) { return i.ID === d.ItemID; });
        var itemName = itemData ? itemData.Name : "Item #" + d.ItemID;
        details += "<div class='mb-2'><strong>" + itemName + "</strong><br>Qty: " + d.Qty + "</div>";
      });
    }
    var supplierName = p.Supplier ? p.Supplier.Name : "-";
    var date = new Date(p.Date).toLocaleString();
    var total = p.GrandTotal ? p.GrandTotal.toLocaleString() : 0;

    Swal.fire({
      title: "Purchase #" + p.ID,
      html: "<div style='text-align: left;'>" +
        "<p><strong>Supplier:</strong> " + supplierName + "</p>" +
        "<p><strong>Date:</strong> " + date + "</p>" +
        "<p><strong>Grand Total:</strong> <span class='text-primary fw-bold'>Rp " + total + "</span></p>" +
        "<hr><h6>Details:</h6>" +
        (details || "<em>Tidak ada detail</em>") +
        "</div>",
      confirmButtonColor: "#0d6efd"
    });
  }).fail(function() {
    Swal.fire({ icon: "error", title: "Gagal Memuat Detail", text: "Terjadi kesalahan saat memuat detail purchase", confirmButtonColor: "#dc3545" });
  });
});

loadPurchases();
