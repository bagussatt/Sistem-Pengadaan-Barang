const BASE_URL = "http://127.0.0.1:3000/api"

function apiRequest(method, endpoint, data, onSuccess, onError) {
  $.ajax({
    url: BASE_URL + endpoint,
    method: method,
    contentType: "application/json",
    headers: {
      Authorization: "Bearer " + localStorage.getItem("token")
    },
    data: data ? JSON.stringify(data) : null,
    success: onSuccess,
    error: function (xhr) {
      let msg = xhr.responseJSON?.message || "Terjadi kesalahan server"
      if (onError) onError(msg)
    }
  })
}

function apiGet(endpoint, cb, err) {
  apiRequest("GET", endpoint, null, cb, err)
}

function apiPost(endpoint, data, cb, err) {
  apiRequest("POST", endpoint, data, cb, err)
}
