public class Oop {
    static class Circle {
        int radius;
        Circle(int r) { this.radius = r; }
        int area() { return 314 * radius * radius / 100; }
        int perimeter() { return 628 * radius / 100; }
    }

    static class Rectangle {
        int width, height;
        Rectangle(int w, int h) { this.width = w; this.height = h; }
        int area() { return width * height; }
        int perimeter() { return 2 * (width + height); }
    }

    public static void main(String[] args) {
        int N = 100000;
        long totalArea = 0;
        long totalPerimeter = 0;

        for (int i = 0; i < N; i++) {
            int r = 1 + i % 100;

            if (i % 3 == 0) {
                Circle c = new Circle(r);
                totalArea += c.area();
                totalPerimeter += c.perimeter();
            } else if (i % 3 == 1) {
                Rectangle rect = new Rectangle(r, r * 2);
                totalArea += rect.area();
                totalPerimeter += rect.perimeter();
            } else {
                Rectangle sq = new Rectangle(r, r);
                totalArea += sq.area();
                totalPerimeter += sq.perimeter();
            }
        }

        System.out.println(totalArea);
        System.out.println(totalPerimeter);
    }
}
