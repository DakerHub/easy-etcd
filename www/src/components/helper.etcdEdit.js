export function listToTree(list, pathKey = "path", nodeFormatter = a => a) {
  list.forEach(el => {
    el.path = el[pathKey];
  });
  let uid = 0;
  const tree = [];

  for (let i = 0; i < list.length; i++) {
    const element = list[i];
    insertNode(tree, element, "/");
  }

  // 挂载节点
  function insertNode(list, node, parentPath) {
    const restPath = node.path.replace(parentPath, "");
    const isLeaf = !restPath.includes("/");

    // 叶子节点，直接添加
    if (isLeaf) {
      const newNode = Object.assign({}, node);
      newNode.id = uid++;
      newNode.name = restPath;
      newNode.isLeaf = true;
      list.push(nodeFormatter(newNode));
      return;
    }

    const curNodePath = restPath.split("/")[0];
    const curFullPath = parentPath + curNodePath + "/";

    const existDir = list.find(item => item.path === curFullPath);

    // 非叶子节点，但是当前目录已存在
    if (existDir) {
      insertNode(existDir.children, node, curFullPath);
      return;
    }

    // 非叶子节点，目录不存在，需新建目录并往下加节点
    const newDir = {
      path: curFullPath,
      name: curNodePath,
      id: uid++,
      children: []
    };
    insertNode(newDir.children, node, curFullPath);
    list.push(nodeFormatter(newDir));
  }
  return tree;
}
