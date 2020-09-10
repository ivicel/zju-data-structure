package ch04;

import java.util.LinkedList;
import java.util.Queue;

public class AvlTree<T extends Comparable<T>> {
    /**
     * 根结点
     */
    private Node<T> root;

    public Node<T> getRoot() {
        return root;
    }

    public void setRoot(Node<T> root) {
        this.root = root;
    }

    public static class Node<T extends Comparable<T>> implements Comparable<Node<T>> {
        /**
         * 结点高度
         */
        private int height = 1;
        /**
         * 存储的数据
         */
        private T value;
        /**
         * 左孩子
         */
        private Node<T> left;
        /**
         * 右孩子
         */
        private Node<T> right;

        public Node(T value) {
            this.value = value;
        }

        @Override
        public int compareTo(Node<T> other) {
            return value.compareTo(other.value);
        }

        @Override
        public String toString() {
            return "Node{" + value + "}";
        }
    }

    private int getHeight(Node<T> node) {
        return node == null ? 0 : node.height;
    }

    /**
     * 插入新结点
     *
     * @param t 结点数据
     */
    public void insert(T t) {
        Node<T> newNode = new Node<>(t);
        root = insert(root, newNode);
    }

    private Node<T> insert(Node<T> parent, Node<T> newNode) {
        if (parent == null) {
            return newNode;
        }

        // 插入到左孩子还是右孩子
        if (parent.compareTo(newNode) > 0) {
            parent.left = insert(parent.left, newNode);
        } else {
            parent.right = insert(parent.right, newNode);
        }

        // 重新计算受影响的父结点深度
        calculateHeight(parent);

        // 父结点是否需要重新平衡
        if (Math.abs(getHeight(parent.left) - getHeight(parent.right)) > 1) {
            parent = makeBalance(parent);
            calculateHeight(parent);
        }

        return parent;
    }

    private void calculateHeight(Node<T> parent) {
        parent.height = Math.max(getHeight(parent.left), getHeight(parent.right)) + 1;
    }

    /**
     * 重新平衡树
     *
     * @param parent 父结点
     * @return 整平衡后的新的父结点
     */
    private Node<T> makeBalance(Node<T> parent) {
        Node<T> node1 = parent;
        Node<T> node2;
        if (getHeight(parent.left) > getHeight(parent.right)) {
            if (getHeight(parent.left.left) >= getHeight(parent.left.right)) { // LL 失衡
                node1 = parent.left;
                parent.left = node1.right;
                node1.right = parent;
            } else if (getHeight(parent.left.left) < getHeight(parent.left.right)) { // LR 失衡
                node1 = parent.left.right;
                node2 = parent.left;

                node2.right = node1.left;
                node1.left = node2;
                parent.left = node1.right;
                node1.right = parent;

                calculateHeight(node2);
                calculateHeight(parent);
            }
        } else if (getHeight(parent.left) < getHeight(parent.right)) { //
            if (getHeight(parent.right.right) >= getHeight(parent.right.left)) { // RR 失衡
                node1 = parent.right;
                parent.right = node1.left;
                node1.left = parent;
            } else if (getHeight(parent.right.right) < getHeight(parent.right.left)) { // RL 失衡
                node1 = parent.right.left;
                node2 = parent.right;

                node2.left = node1.right;
                node1.right = node2;
                parent.right = node1.left;
                node1.left = parent;

                calculateHeight(node2);
                calculateHeight(parent);
            }
        }

        return node1;
    }

    @Override
    public String toString() {
        StringBuilder builder = new StringBuilder();
        builder.append("[");
        if (root != null) {
            boolean flag = false;

            Queue<Node<T>> queue = new LinkedList<>();
            queue.offer(root);
            while (!queue.isEmpty()) {
                Node<T> node = queue.poll();
                if (flag) {
                    builder.append(", ");
                } else {
                    flag = true;
                }
                builder.append(node);

                if (node.left != null) {
                    queue.offer(node.left);
                }

                if (node.right != null) {
                    queue.offer(node.right);
                }
            }
        }

        builder.append("]");

        return builder.toString();
    }

    public static void main(String[] args) {
        AvlTree<Integer> avlTree = new AvlTree<>();
        int[] arr = new int[] { 16, 3, 7, 11, 9, 26, 18, 14, 15 };
        for (int i : arr) {
            avlTree.insert(i);
        }

        System.out.println(avlTree);
    }
}
