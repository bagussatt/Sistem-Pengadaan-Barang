apiRequest("GET", "/items").done(items => {
  let html = `
    <tr>
      <th>ID</th><th>Name</th><th>Stock</th><th>Price</th>
    </tr>`;
  items.forEach(i => {
    html += `
      <tr>
        <td>${i.id}</td>
        <td>${i.name}</td>
        <td>${i.stock}</td>
        <td>${i.price}</td>
      </tr>`;
  });
  $("#itemTable").html(html);
});
