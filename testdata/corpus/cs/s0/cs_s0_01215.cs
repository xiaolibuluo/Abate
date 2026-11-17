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

    public struct Point11400
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 81;
    }

    public static class Grid11400
    {
        public static Point11400 Make(int left, int right)
        {
            return new Point11400 { X = left, Y = right };
        }
    }
    public static int Scale11401<T>(T value) where T : struct
    {
        int number = 88;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 9; }
        return number;
    }
    public async Task<int> Aggregate11402(IEnumerable<int> source)
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
