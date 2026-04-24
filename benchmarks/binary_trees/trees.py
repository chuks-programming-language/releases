import sys
sys.setrecursionlimit(200000)

class TreeNode:
    __slots__ = ('left', 'right', 'value')
    def __init__(self, value):
        self.value = value
        self.left = None
        self.right = None
    def check(self):
        if self.left is None:
            return self.value
        return self.value + self.left.check() - self.right.check()

def build_tree(depth):
    if depth == 0:
        return TreeNode(1)
    node = TreeNode(depth)
    node.left = build_tree(depth - 1)
    node.right = build_tree(depth - 1)
    return node

max_depth = 16
tree = build_tree(max_depth)
print(tree.check())

total = 0
iterations = 1000
for i in range(iterations):
    t = build_tree(10)
    total += t.check()
print(total)
