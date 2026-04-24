import java.util.HashMap;

public class Maps {
    public static void main(String[] args) {
        int N = 100000;

        HashMap<String, Integer> m = new HashMap<>();
        for (int i = 0; i < N; i++) {
            m.put("key_" + i, i * 2);
        }
        System.out.println(m.size());

        int found = 0;
        for (int i = 0; i < N; i++) {
            if (m.containsKey("key_" + i)) found++;
        }
        System.out.println(found);

        long sum = 0;
        for (int v : m.values()) {
            sum += v;
        }
        System.out.println(sum);

        for (int i = 0; i < N; i += 2) {
            m.remove("key_" + i);
        }
        System.out.println(m.size());
    }
}
