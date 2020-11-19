import clonedeep from "lodash.clonedeep";

let uid = 0;

export function listToTree(list, pathKey = "path", nodeFormatter = a => a) {
  const _list = clonedeep(list);
  _list.forEach(el => {
    el.path = el[pathKey];
    if (!el.path.startsWith("/")) {
      el.path = "/" + el.path;
    }
  });
  const tree = [];

  for (let i = 0; i < _list.length; i++) {
    const element = _list[i];
    insertNode(tree, element, "/", nodeFormatter);
  }

  return tree;
}

export function insertNode(list, node, parentPath, nodeFormatter = a => a) {
  const restPath = node.path.replace(parentPath, "");
  const isLeaf = !restPath.includes("/");

  // 叶子节点，直接添加
  if (isLeaf) {
    const newNode = Object.assign(
      {
        id: uid++,
        name: restPath,
        isLeaf: true
      },
      node
    );

    list.push(nodeFormatter(newNode));
    return;
  }

  const curNodePath = restPath.split("/")[0];
  const curFullPath = parentPath + curNodePath + "/";

  const existDir = list.find(item => item.path === curFullPath);

  // 非叶子节点，但是当前目录已存在
  if (existDir) {
    insertNode(existDir.children, node, curFullPath, nodeFormatter);
    return;
  }

  // 非叶子节点，目录不存在，需新建目录并往下加节点
  const newDir = {
    path: curFullPath,
    name: curNodePath,
    id: uid++,
    children: [],
    isLeaf: false
  };
  list.push(nodeFormatter(newDir));
  insertNode(newDir.children, node, curFullPath, nodeFormatter);
}

export function removeTreeNode(tree, key) {
  if (!Array.isArray(tree)) return;

  for (let i = 0; i < tree.length; i++) {
    const node = tree[i];

    if (node.path === key) {
      tree.splice(i, 1);
      return;
    }

    if (key.startsWith(node.path)) {
      removeTreeNode(node.children, key);
      if (!node.children.length) {
        tree.splice(i, 1);
      }
      return;
    }
  }
}
