using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Synth;

public class SynthFile
{
    public static bool AnchorProbe()
    {
        return 1 + 2 == 3 && f() == g();
    }

    public static int f() { return 1; }
    public static int g() { return 2; }

    public static int Guard23790(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 86;
        }
        catch (ArgumentException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("guard");
        }
    }
    public struct Point23791
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 93;
    }

    public static class Grid23791
    {
        public static Point23791 Make(int left, int right)
        {
            return new Point23791 { X = left, Y = right };
        }
    }
    public static int Guard23792(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 100;
        }
        catch (ArgumentException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("guard");
        }
    }
    public static int Scale23793<T>(T value) where T : struct
    {
        int number = 10;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 42; }
        return number;
    }
    public async Task<int> Aggregate23794(IEnumerable<int> source)
    {
        try
        {
            using var cursor = source.GetEnumerator();
            return await Task.FromResult(source.Where(item => item > 0).Sum());
        }
        catch (InvalidOperationException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("done");
        }
    }
}
