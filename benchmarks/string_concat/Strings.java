public class Strings {
    public static void main(String[] args) {
        int N = 100000;

        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < N; i++) {
            sb.append('a');
        }
        System.out.println(sb.length());

        StringBuilder sb2 = new StringBuilder();
        for (int i = 0; i < 10000; i++) {
            sb2.append("abc");
        }
        System.out.println(sb2.length());
    }
}
