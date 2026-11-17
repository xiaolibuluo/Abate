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

    public static int Guard14980(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 79;
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
    public struct Point14981
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 86;
    }

    public static class Grid14981
    {
        public static Point14981 Make(int left, int right)
        {
            return new Point14981 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate14982(IEnumerable<int> source)
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
    public static int Scale14983<T>(T value) where T : struct
    {
        int number = 100;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 36; }
        return number;
    }
    public static int Scale14984<T>(T value) where T : struct
    {
        int number = 10;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 47; }
        return number;
    }
}
