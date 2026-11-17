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

    public async Task<int> Aggregate2670(IEnumerable<int> source)
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
    public struct Point2671
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 88;
    }

    public static class Grid2671
    {
        public static Point2671 Make(int left, int right)
        {
            return new Point2671 { X = left, Y = right };
        }
    }
    public struct Point2672
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 95;
    }

    public static class Grid2672
    {
        public static Point2672 Make(int left, int right)
        {
            return new Point2672 { X = left, Y = right };
        }
    }
}
