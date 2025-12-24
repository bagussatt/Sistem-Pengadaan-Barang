let cart = [];

$(document).ready(function () {

  apiRequest("GET", "/api/suppliers")
    .done(data => {
      data.forEach(s =>
        $("#supplier").append(`<option value="${s.id}">${s.name}</option>`)
      );
    });

  apiRequest("GET", "/api/items")
    .done(data => {
      data.forEach(i =>
        $("#item").append(`<option value="${i.id}">${i.name}</option>`)
      );
    });
});

$("#addItem").on("click", function () {
  cart.push({
    item_id: parseInt($("#item").val()),
    qty: parseInt($("#qty").val())
  });
  renderCart();
});

$("#cartTable").on("click", ".remove", function () {
  cart.splice($(this).data("index"), 1);
  renderCart();
});

function renderCart() {
  $("#cartTable").empty();
  cart.forEach((c, i) => {
    $("#cartTable").append(`
      <tr>
        <td>${c.item_id}</td>
        <td>${c.qty}</td>
        <td><button class="btn btn-danger btn-sm remove" data-index="${i}">Hapus</button></td>
      </tr>
    `);
  });
}

$("#submitOrder").on("click", function () {
  apiRequest("POST", "/api/purchasings", {
    supplier_id: parseInt($("#supplier").val()),
    items: cart
  })
  .done(() => {
    alert("Purchase berhasil");
    cart = [];
    renderCart();
  })
  .fail(err => showError(err));
});
