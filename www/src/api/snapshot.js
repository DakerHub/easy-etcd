import request from "./request";

export function getSnapshots(node) {
  return request.post("/snapshot/query", {
    nodes: [node]
  });
}

export function createSnapshot(node) {
  return request.post("/snapshot/create", {
    nodes: [node]
  });
}

export function deleteSnapshot(node, snapshotID) {
  return request.post("/snapshot/delete", {
    nodes: [node],
    snapshotID
  });
}

export function restore(node, snapshotID) {
  return request.post("/restore", {
    nodes: [node],
    snapshotID,
    clear: true
  });
}
