import request from "./request";

export function login(username, password) {
  return request.post("/login", {
    username,
    password
  });
}

export function testConnect(node) {
  return request.post("/connect", {
    nodes: [node]
  });
}
