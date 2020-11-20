import request from "./request";

export function testConnect(node) {
  return request.post("/connect", {
    nodes: [node]
  });
}
