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

    public struct Point10980
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 20;
    }

    public static class Grid10980
    {
        public static Point10980 Make(int left, int right)
        {
            return new Point10980 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate10981(IEnumerable<int> source)
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
    public struct Point10982
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 34;
    }

    public static class Grid10982
    {
        public static Point10982 Make(int left, int right)
        {
            return new Point10982 { X = left, Y = right };
        }
    }
}
