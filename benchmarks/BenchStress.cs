// ============================================================
// C# PERFORMANCE BENCHMARK SUITE
// All times in milliseconds
// ============================================================

using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Text;

class BenchStress
{
    static long Fib(int n)
    {
        long a = 0, b = 1;
        for (int i = 0; i < n; i++)
        {
            long tmp = a + b;
            a = b;
            b = tmp;
        }
        return a;
    }

    static long MatmulSim(int size)
    {
        long sum = 0;
        for (int i = 0; i < size; i++)
            for (int j = 0; j < size; j++)
                for (int k = 0; k < size; k++)
                    sum = sum + (i * k) - (j * k) + 1;
        return sum;
    }

    static int CountPrimes(int limit)
    {
        int count = 0;
        for (int n = 2; n < limit; n++)
        {
            bool isPrime = true;
            for (int d = 2; d * d <= n; d++)
            {
                if (n % d == 0) { isPrime = false; break; }
            }
            if (isPrime) count++;
        }
        return count;
    }

    static long ArrayStress(int n)
    {
        var arr = new List<long>();
        for (int i = 0; i < n; i++)
            arr.Add((long)i * i);
        long sum = 0;
        for (int i = 0; i < arr.Count; i++)
            sum += arr[i];
        return sum;
    }

    static long MapStress(int n)
    {
        var m = new Dictionary<string, long>();
        for (int i = 0; i < n; i++)
            m["key"] = (long)i * i;
        long sum = 0;
        foreach (var key in m.Keys)
            sum += m[key];
        return sum;
    }

    class Vector
    {
        public long X, Y;
        public Vector(long x, long y) { X = x; Y = y; }
        public Vector Add(Vector other) => new Vector(X + other.X, Y + other.Y);
        public long MagnitudeSquared() => X * X + Y * Y;
    }

    static long OopStress(int n)
    {
        long total = 0;
        for (int i = 0; i < n; i++)
        {
            var v1 = new Vector(i, i + 1);
            var v2 = new Vector(i + 2, i + 3);
            var v3 = v1.Add(v2);
            total += v3.MagnitudeSquared();
        }
        return total;
    }

    class Shape
    {
        public string Name;
        public Shape(string name) { Name = name; }
        public virtual long Area() => 0;
    }

    class Rectangle : Shape
    {
        public long W, H;
        public Rectangle(long w, long h) : base("rect") { W = w; H = h; }
        public override long Area() => W * H;
    }

    static long InheritanceStress(int n)
    {
        long total = 0;
        for (int i = 0; i < n; i++)
        {
            var r = new Rectangle(i, i + 1);
            total += r.Area();
        }
        return total;
    }

    static Func<int> MakeCounter()
    {
        int count = 0;
        return () => { count++; return count; };
    }

    static int ClosureStress(int n)
    {
        var counter = MakeCounter();
        int result = 0;
        for (int i = 0; i < n; i++)
            result = counter();
        return result;
    }

    static int StringStress(int n)
    {
        var sb = new StringBuilder();
        for (int i = 0; i < n; i++)
            sb.Append("x");
        return sb.Length;
    }

    static int FibRecursive(int n)
    {
        if (n <= 1) return n;
        return FibRecursive(n - 1) + FibRecursive(n - 2);
    }

    static long Ms(Stopwatch sw) => sw.ElapsedMilliseconds;

    static void Main()
    {
        Console.WriteLine("=== C# PERFORMANCE BENCHMARK SUITE ===");

        var sw = Stopwatch.StartNew();
        var t0 = sw.ElapsedMilliseconds;

        Fib(1000000);
        var t1 = sw.ElapsedMilliseconds;
        Console.WriteLine("1_fib_iter_1M");
        Console.WriteLine(t1 - t0);

        var t2 = sw.ElapsedMilliseconds;
        MatmulSim(100);
        var t3 = sw.ElapsedMilliseconds;
        Console.WriteLine("2_nested_loops_1M");
        Console.WriteLine(t3 - t2);

        var t4 = sw.ElapsedMilliseconds;
        CountPrimes(50000);
        var t5 = sw.ElapsedMilliseconds;
        Console.WriteLine("3_primes_50K");
        Console.WriteLine(t5 - t4);

        var t6 = sw.ElapsedMilliseconds;
        ArrayStress(100000);
        var t7 = sw.ElapsedMilliseconds;
        Console.WriteLine("4_array_100K");
        Console.WriteLine(t7 - t6);

        var t8 = sw.ElapsedMilliseconds;
        MapStress(50000);
        var t9 = sw.ElapsedMilliseconds;
        Console.WriteLine("5_map_50K");
        Console.WriteLine(t9 - t8);

        var t10 = sw.ElapsedMilliseconds;
        OopStress(100000);
        var t11 = sw.ElapsedMilliseconds;
        Console.WriteLine("6_oop_100K");
        Console.WriteLine(t11 - t10);

        var t12 = sw.ElapsedMilliseconds;
        InheritanceStress(100000);
        var t13 = sw.ElapsedMilliseconds;
        Console.WriteLine("7_inherit_100K");
        Console.WriteLine(t13 - t12);

        var t14 = sw.ElapsedMilliseconds;
        ClosureStress(1000000);
        var t15 = sw.ElapsedMilliseconds;
        Console.WriteLine("8_closure_1M");
        Console.WriteLine(t15 - t14);

        var t16 = sw.ElapsedMilliseconds;
        StringStress(10000);
        var t17 = sw.ElapsedMilliseconds;
        Console.WriteLine("9_string_10K");
        Console.WriteLine(t17 - t16);

        var t18 = sw.ElapsedMilliseconds;
        FibRecursive(30);
        var t19 = sw.ElapsedMilliseconds;
        Console.WriteLine("10_fib_rec_30");
        Console.WriteLine(t19 - t18);

        Console.WriteLine("TOTAL");
        Console.WriteLine(t19 - t0);
    }
}
