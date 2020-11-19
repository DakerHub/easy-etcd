import request from "./request";

export function getAllKvs(node) {
  return request.post("/kv/query", {
    nodes: [node]
  });
}

export function deleteKv(node, key) {
  return request.post("/kv/delete", {
    nodes: [node],
    key
  });
}

export function putKv(node, key, value) {
  return request.post("/kv/put", {
    nodes: [node],
    key,
    value
  });
}
