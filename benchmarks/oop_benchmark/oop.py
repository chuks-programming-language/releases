class Circle:
    def __init__(self, radius):
        self.radius = radius
    def area(self):
        return 314 * self.radius * self.radius // 100
    def perimeter(self):
        return 628 * self.radius // 100

class Rectangle:
    def __init__(self, width, height):
        self.width = width
        self.height = height
    def area(self):
        return self.width * self.height
    def perimeter(self):
        return 2 * (self.width + self.height)

N = 100000
total_area = 0
total_perimeter = 0

for i in range(N):
    r = 1 + i % 100

    if i % 3 == 0:
        c = Circle(r)
        total_area += c.area()
        total_perimeter += c.perimeter()
    elif i % 3 == 1:
        rect = Rectangle(r, r * 2)
        total_area += rect.area()
        total_perimeter += rect.perimeter()
    else:
        sq = Rectangle(r, r)
        total_area += sq.area()
        total_perimeter += sq.perimeter()

print(total_area)
print(total_perimeter)
