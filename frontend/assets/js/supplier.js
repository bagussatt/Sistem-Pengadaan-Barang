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
