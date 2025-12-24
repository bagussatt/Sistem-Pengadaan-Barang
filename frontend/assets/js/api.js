const BASE_URL = "http://127.0.0.1:3000";

function apiRequest(method, url, data = null) {
  return $.ajax({
    method,
    url: BASE_URL + url,
    contentType: "application/json",
    data: data ? JSON.stringify(data) : null,
    headers: {
      Authorization: "Bearer " + localStorage.getItem("token")
    }
  });
}

function showError(err) {
  Swal.fire({
    icon: "error",
    title: "Error",
    text: err.responseJSON?.message || "Server error"
  });
}
