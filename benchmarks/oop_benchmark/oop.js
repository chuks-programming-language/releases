class Circle {
  constructor(radius) {
    this.radius = radius;
  }
  area() {
    return Math.trunc((314 * this.radius * this.radius) / 100);
  }
  perimeter() {
    return Math.trunc((628 * this.radius) / 100);
  }
}

class Rectangle {
  constructor(width, height) {
    this.width = width;
    this.height = height;
  }
  area() {
    return this.width * this.height;
  }
  perimeter() {
    return 2 * (this.width + this.height);
  }
}

const N = 100000;
let totalArea = 0;
let totalPerimeter = 0;

for (let i = 0; i < N; i++) {
  const r = 1 + (i % 100);

  if (i % 3 === 0) {
    const c = new Circle(r);
    totalArea += c.area();
    totalPerimeter += c.perimeter();
  } else if (i % 3 === 1) {
    const rect = new Rectangle(r, r * 2);
    totalArea += rect.area();
    totalPerimeter += rect.perimeter();
  } else {
    const sq = new Rectangle(r, r);
    totalArea += sq.area();
    totalPerimeter += sq.perimeter();
  }
}

console.log(totalArea);
console.log(totalPerimeter);
