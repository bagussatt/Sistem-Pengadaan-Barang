const BASE_URL = "http://127.0.0.1:3000/api";

function apiPost(url, data) {
  return $.ajax({
    url: BASE_URL + url,
    method: "POST",
    contentType: "application/json",
    data: JSON.stringify(data),
  });
}
function authHeader() {
  return {
    Authorization: "Bearer " + localStorage.getItem("token")
  };
}

function apiGet(url) {
  return $.ajax({
    url: BASE_URL + url,
    method: "GET",
    headers: authHeader()
  });
}

function apiDelete(url) {
  return $.ajax({
    url: BASE_URL + url,
    method: "DELETE",
    headers: authHeader()
  });
}
