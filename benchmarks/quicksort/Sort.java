public class Sort {
    static int[] arr;

    static void quicksort(int lo, int hi) {
        if (lo >= hi) return;
        int pivot = arr[hi];
        int i = lo;
        for (int j = lo; j < hi; j++) {
            if (arr[j] <= pivot) {
                int tmp = arr[i]; arr[i] = arr[j]; arr[j] = tmp;
                i++;
            }
        }
        int tmp = arr[i]; arr[i] = arr[hi]; arr[hi] = tmp;
        quicksort(lo, i - 1);
        quicksort(i + 1, hi);
    }

    public static void main(String[] args) {
        int N = 100000;
        arr = new int[N];
        int seed = 42;
        for (int i = 0; i < N; i++) {
            seed = (int)((((long)seed * 1103515245L + 12345L) % 2147483648L));
            arr[i] = seed % 1000000;
        }

        quicksort(0, N - 1);

        int sorted = 1;
        for (int i = 1; i < N; i++) {
            if (arr[i] < arr[i - 1]) sorted = 0;
        }
        System.out.println(sorted);
        System.out.println(arr[0]);
        System.out.println(arr[N - 1]);
    }
}
