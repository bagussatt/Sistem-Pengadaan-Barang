$(document).ready(function () {
  apiRequest("GET", "/api/items")
    .done(items => {
      items.forEach(i => {
        $("#itemsTable").append(`
          <tr>
            <td>${i.name}</td>
            <td>${i.stock}</td>
            <td>${i.price}</td>
          </tr>
        `);
      });
    })
    .fail(err => showError(err));
});
