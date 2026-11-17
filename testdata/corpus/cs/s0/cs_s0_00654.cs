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

    public struct Point6100
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 78;
    }

    public static class Grid6100
    {
        public static Point6100 Make(int left, int right)
        {
            return new Point6100 { X = left, Y = right };
        }
    }
    public struct Point6101
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 85;
    }

    public static class Grid6101
    {
        public static Point6101 Make(int left, int right)
        {
            return new Point6101 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate6102(IEnumerable<int> source)
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
