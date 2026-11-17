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

    public struct Point10460
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 23;
    }

    public static class Grid10460
    {
        public static Point10460 Make(int left, int right)
        {
            return new Point10460 { X = left, Y = right };
        }
    }
    public struct Point10461
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 30;
    }

    public static class Grid10461
    {
        public static Point10461 Make(int left, int right)
        {
            return new Point10461 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate10462(IEnumerable<int> source)
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
