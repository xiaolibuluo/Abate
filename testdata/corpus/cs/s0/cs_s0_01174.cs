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

    public struct Point11000
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 46;
    }

    public static class Grid11000
    {
        public static Point11000 Make(int left, int right)
        {
            return new Point11000 { X = left, Y = right };
        }
    }
    public struct Point11001
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 53;
    }

    public static class Grid11001
    {
        public static Point11001 Make(int left, int right)
        {
            return new Point11001 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate11002(IEnumerable<int> source)
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
