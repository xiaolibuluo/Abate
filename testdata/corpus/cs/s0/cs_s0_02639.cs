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

    public static int Guard24690(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 92;
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
    public struct Point24691
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 99;
    }

    public static class Grid24691
    {
        public static Point24691 Make(int left, int right)
        {
            return new Point24691 { X = left, Y = right };
        }
    }
    public struct Point24692
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 9;
    }

    public static class Grid24692
    {
        public static Point24692 Make(int left, int right)
        {
            return new Point24692 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate24693(IEnumerable<int> source)
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
