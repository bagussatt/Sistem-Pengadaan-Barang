let cart = []

$(document).ready(function () {
  loadSuppliers()
  loadItems()
})

function loadSuppliers() {
  $("#supplier").empty().append(`<option value="">-- Pilih Supplier --</option>`)
  apiGet("/suppliers", function (data) {
    data.forEach(s => {
      $("#supplier").append(`<option value="${s.id}">${s.name}</option>`)
    })
  })
}

function loadItems() {
  $("#item").empty().append(`<option value="">-- Pilih Item --</option>`)
  apiGet("/items", function (data) {
    data.forEach(i => {
      $("#item").append(`
        <option value="${i.id}" data-name="${i.name}" data-stock="${i.stock}">
          ${i.name} (Stock: ${i.stock})
        </option>
      `)
    })
  })
}

$("#addItem").on("click", function () {
  const id = $("#item").val()
  const name = $("#item option:selected").data("name")
  const stock = $("#item option:selected").data("stock")
  const qty = parseInt($("#qty").val())

  if (!id || !qty || qty <= 0)
    return Swal.fire("Error", "Item dan Qty wajib diisi", "error")

  if (qty > stock)
    return Swal.fire("Error", "Qty melebihi stok", "error")

  cart.push({ item_id: parseInt(id), item_name: name, qty })
  $("#qty").val("")
  renderCart()
})

function renderCart() {
  $("#cartTable").empty()
  cart.forEach((i, idx) => {
    $("#cartTable").append(`
      <tr>
        <td>${i.item_name}</td>
        <td>${i.qty}</td>
        <td>
          <button class="btn btn-danger btn-sm delete-item" data-index="${idx}">
            Hapus
          </button>
        </td>
      </tr>
    `)
  })
}

$("#cartTable").on("click", ".delete-item", function () {
  cart.splice($(this).data("index"), 1)
  renderCart()
})

$("#submitOrder").on("click", function () {
  if (!$("#supplier").val())
    return Swal.fire("Error", "Supplier wajib dipilih", "error")

  if (cart.length === 0)
    return Swal.fire("Error", "Keranjang kosong", "error")

  apiPost("/purchasings", {
    supplier_id: parseInt($("#supplier").val()),
    items: cart.map(i => ({ item_id: i.item_id, qty: i.qty }))
  }, function () {
    Swal.fire("Sukses", "Transaksi berhasil disimpan", "success")
    cart = []
    renderCart()
  }, function (msg) {
    Swal.fire("Gagal", msg, "error")
  })
})

$("#saveSupplier").on("click", function () {
  apiPost("/suppliers", {
    name: $("#supName").val(),
    email: $("#supEmail").val(),
    address: $("#supAddress").val()
  }, function () {
    Swal.fire("Sukses", "Supplier ditambahkan", "success")
    $("#supName,#supEmail,#supAddress").val("")
    bootstrap.Modal.getInstance(
      document.getElementById("supplierModal")
    ).hide()
    loadSuppliers()
  }, function (msg) {
    Swal.fire("Error", msg, "error")
  })
})
