apiRequest("GET", "/suppliers").done(data => {
  let html = `
    <tr>
      <th>ID</th><th>Name</th><th>Email</th><th>Address</th>
    </tr>`;
  data.forEach(s => {
    html += `
      <tr>
        <td>${s.id}</td>
        <td>${s.name}</td>
        <td>${s.email || "-"}</td>
        <td>${s.address || "-"}</td>
      </tr>`;
  });
  $("#supplierTable").html(html);
});
