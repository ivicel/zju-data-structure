package ch05;

@SuppressWarnings("unchecked")
public class PriorityQueue<T extends Comparable<T>> {
    /**
     * 当前存的节点容量大小
     */
    private int size;

    /**
     * 最大的容量大小
     */
    private int maxSize;

    /**
     * 数据存储地方
     */
    private Object[] values;

    public PriorityQueue(int maxSize) {
        this.maxSize = maxSize;
        this.size = 0;
        this.values = new Object[maxSize];
    }

    private T get(int i) {
        return (T) values[i];
    }

    /**
     * 获取最大值
     *
     * @return
     */
    public T peek() {
        return (T) values[1];
    }

    /**
     * 插入新结点
     */
    public void push(T t) {
        if (size == maxSize) {
            throw new UnsupportedOperationException("堆已满");
        }

        // 将新结点插到树最后的位置
        values[0] = values[++size] = t;
        // 上滤调整
        percolateUp();
    }

    /**
     * 上滤
     */
    protected void percolateUp() {
        for (int i = size; get(i / 2).compareTo(get(0)) < 0; i = i / 2) {
            Object node = values[i / 2];
            values[i / 2] = values[i];
            values[i] = node;
        }
    }

    /**
     * 删除最大结点
     */
    public T pop() {
        if (size == 0) {
            throw new UnsupportedOperationException("无法删除空树");
        }

        Object node = values[1];
        values[0] = values[1] = values[size--];
        percolateDown(1);

        return (T) node;
    }

    /**
     * 下滤
     */
    protected void percolateDown(int parent) {
        // 如果左孩子都不存在的话, 那就是叶子结点了, 因为这是一棵完全二叉树
        while (parent * 2 <= size) {
            int pos = parent * 2;
            int rightPos = parent * 2 + 1;

            // 如果右子树存在并且右 > 左, 那 pos = 右, 否则 pos = 左
            if (rightPos <= size && get(rightPos).compareTo(get(pos)) > 0) {
                pos = rightPos;
            }

            // 比较孩子结点和父结点的大小, pos 可能是左也可能是右
            // 不大于则说明最大结点就是父结点
            if (get(parent).compareTo(get(pos)) < 0) {
                Object tmp = values[parent];
                values[parent] = values[pos];
                values[pos] = tmp;
                // 转到下一层
                parent = pos;
            } else {
                // 父结点是最大的就退出比较, 否则交换位置
                break;
            }
        }
    }

    public static <T extends Comparable<T>> PriorityQueue<T> createQueue(T... arr) {
        PriorityQueue<T> queue = new PriorityQueue<>(arr.length * 2);
        queue.size = arr.length;
        System.arraycopy(arr, 0, queue.values, 1, arr.length);

        for (int i = queue.size / 2; i > 0; i--) {
            queue.percolateDown(i);
        }

        return queue;
    }

    @Override
    public String toString() {
        StringBuilder builder = new StringBuilder(size);
        builder.append("[");
        boolean flag = false;
        for (int i = 1; i <= size; i++) {
            if (flag) {
                builder.append(", ");
            } else {
                flag = true;
            }
            builder.append(values[i]);
        }

        return builder.append("]").toString();
    }

    public static void main(String[] args) {
        Integer[] arr = { 79, 66, 16, 83, 30, 19, 68, 55, 91, 72, 49, 9 };
        PriorityQueue<Integer> queue = new PriorityQueue<>(20);

        System.out.println("插入数据演示:");
        for (int value : arr) {
            System.out.println("插入新数据: " + value);
            queue.push(value);
        }

        System.out.println("生成的堆: " + queue);
        System.out.println("---------------------------------------\n");

        System.out.println("删除数据演示:");
        while (queue.size > 0) {
            Integer value = queue.pop();
            System.out.println("删除数据: " + value);
        }

        System.out.println("---------------------------------------\n");

        queue = PriorityQueue.createQueue(arr);
        System.out.println("O(N) 时间复杂度生成一个新堆: " + queue);
    }
}
