let cart = [];


apiRequest("GET", "/suppliers").done(data => {
  data.forEach(s => {
    $("#supplier").append(`<option value="${s.id}">${s.name}</option>`);
  });
});


apiRequest("GET", "/items").done(data => {
  data.forEach(i => {
    $("#item").append(`<option value="${i.id}">${i.name}</option>`);
  });
});


$("#addItem").click(() => {
  cart.push({
    item_id: Number($("#item").val()),
    qty: Number($("#qty").val())
  });
  renderCart();
});

function renderCart() {
  let html = "";
  cart.forEach((c, i) => {
    html += `
      <tr>
        <td>${c.item_id}</td>
        <td>${c.qty}</td>
        <td><button class="btn btn-danger btn-sm remove" data-i="${i}">Hapus</button></td>
      </tr>`;
  });
  $("#cartTable").html(html);
}


$(document).on("click", ".remove", function () {
  cart.splice($(this).data("i"), 1);
  renderCart();
});


$("#submitOrder").click(() => {
  apiRequest("POST", "/purchasings", {
    supplier_id: Number($("#supplier").val()),
    items: cart
  })
  .done(() => {
    Swal.fire("Success", "Purchase berhasil", "success");
    cart = [];
    renderCart();
  })
  .fail(err => {
    Swal.fire("Error", err.responseJSON.message, "error");
  });
});
