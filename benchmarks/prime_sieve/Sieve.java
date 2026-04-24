public class Sieve {
    public static void main(String[] args) {
        int N = 1000000;
        boolean[] sieve = new boolean[N];
        for (int i = 0; i < N; i++) sieve[i] = true;
        sieve[0] = false;
        sieve[1] = false;

        for (int i = 2; i * i < N; i++) {
            if (sieve[i]) {
                for (int j = i * i; j < N; j += i) {
                    sieve[j] = false;
                }
            }
        }

        int count = 0;
        for (int i = 0; i < N; i++) {
            if (sieve[i]) count++;
        }
        System.out.println(count);
    }
}
