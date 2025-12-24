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
