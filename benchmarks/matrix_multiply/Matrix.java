public class Matrix {
    public static void main(String[] args) {
        int N = 200;

        int[] A = new int[N * N];
        int[] B = new int[N * N];
        int[] C = new int[N * N];

        for (int i = 0; i < N * N; i++) {
            A[i] = i % 100;
            B[i] = (i * 3 + 7) % 100;
        }

        for (int i = 0; i < N; i++) {
            for (int j = 0; j < N; j++) {
                int sum = 0;
                for (int k = 0; k < N; k++) {
                    sum += A[i * N + k] * B[k * N + j];
                }
                C[i * N + j] = sum;
            }
        }

        int checksum = 0;
        for (int i = 0; i < N; i++) {
            checksum += C[i * N + i];
        }
        System.out.println(checksum);
    }
}
