const BASE_URL = "http://127.0.0.1:3000/api";

function authHeader() {
  const token = localStorage.getItem("token");
  return token ? { Authorization: "Bearer " + token } : {};
}

function apiRequest({ url, method = "GET", data = null }) {
  return $.ajax({
    url: BASE_URL + url,
    method: method,
    contentType: "application/json",
    headers: authHeader(),
    data: data ? JSON.stringify(data) : null
  });
}

function apiGet(url) {
  return $.ajax({
    url: BASE_URL + url,
    method: "GET",
    headers: authHeader()
  });
}

function apiPost(url, data) {
  return $.ajax({
    url: BASE_URL + url,
    method: "POST",
    contentType: "application/json",
    headers: authHeader(), 
    data: JSON.stringify(data)
  });
}

function apiDelete(url) {
  return $.ajax({
    url: BASE_URL + url,
    method: "DELETE",
    headers: authHeader()
  });
}

function showError(err, fallback = "Terjadi kesalahan") {
  const msg =
    err.responseJSON?.message ||
    err.responseJSON?.error ||
    fallback;

  Swal.fire("Error", msg, "error");
}
