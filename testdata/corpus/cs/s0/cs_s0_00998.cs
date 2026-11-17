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

    public struct Point9360
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 48;
    }

    public static class Grid9360
    {
        public static Point9360 Make(int left, int right)
        {
            return new Point9360 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate9361(IEnumerable<int> source)
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
