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

    public struct Point18740
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 20;
    }

    public static class Grid18740
    {
        public static Point18740 Make(int left, int right)
        {
            return new Point18740 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate18741(IEnumerable<int> source)
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
    public static int Scale18742<T>(T value) where T : struct
    {
        int number = 34;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 34; }
        return number;
    }
}
