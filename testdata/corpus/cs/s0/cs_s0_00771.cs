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

    public struct Point7230
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 92;
    }

    public static class Grid7230
    {
        public static Point7230 Make(int left, int right)
        {
            return new Point7230 { X = left, Y = right };
        }
    }
    public struct Point7231
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 99;
    }

    public static class Grid7231
    {
        public static Point7231 Make(int left, int right)
        {
            return new Point7231 { X = left, Y = right };
        }
    }
    public struct Point7232
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 9;
    }

    public static class Grid7232
    {
        public static Point7232 Make(int left, int right)
        {
            return new Point7232 { X = left, Y = right };
        }
    }
    public struct Point7233
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 16;
    }

    public static class Grid7233
    {
        public static Point7233 Make(int left, int right)
        {
            return new Point7233 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate7234(IEnumerable<int> source)
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
