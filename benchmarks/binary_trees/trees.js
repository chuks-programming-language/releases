class TreeNode {
  constructor(value) {
    this.value = value;
    this.left = null;
    this.right = null;
  }
  check() {
    if (this.left === null) return this.value;
    return this.value + this.left.check() - this.right.check();
  }
}

function buildTree(depth) {
  if (depth === 0) return new TreeNode(1);
  const node = new TreeNode(depth);
  node.left = buildTree(depth - 1);
  node.right = buildTree(depth - 1);
  return node;
}

const maxDepth = 16;
const tree = buildTree(maxDepth);
console.log(tree.check());

let sum = 0;
const iterations = 1000;
for (let i = 0; i < iterations; i++) {
  const t = buildTree(10);
  sum += t.check();
}
console.log(sum);
