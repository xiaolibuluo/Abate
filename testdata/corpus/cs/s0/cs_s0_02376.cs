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

    public struct Point22180
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 30;
    }

    public static class Grid22180
    {
        public static Point22180 Make(int left, int right)
        {
            return new Point22180 { X = left, Y = right };
        }
    }
    public static int Scale22181<T>(T value) where T : struct
    {
        int number = 37;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 52; }
        return number;
    }
    public async Task<int> Aggregate22182(IEnumerable<int> source)
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
    public static int Guard22183(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 51;
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
}
