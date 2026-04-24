public class Trees {
    static class TreeNode {
        TreeNode left, right;
        int value;
        TreeNode(int value) { this.value = value; }
        int check() {
            if (left == null) return value;
            return value + left.check() - right.check();
        }
    }

    static TreeNode buildTree(int depth) {
        if (depth == 0) { TreeNode n = new TreeNode(1); return n; }
        TreeNode node = new TreeNode(depth);
        node.left = buildTree(depth - 1);
        node.right = buildTree(depth - 1);
        return node;
    }

    public static void main(String[] args) {
        int maxDepth = 16;
        TreeNode tree = buildTree(maxDepth);
        System.out.println(tree.check());

        int sum = 0;
        int iterations = 1000;
        for (int i = 0; i < iterations; i++) {
            TreeNode t = buildTree(10);
            sum += t.check();
        }
        System.out.println(sum);
    }
}
