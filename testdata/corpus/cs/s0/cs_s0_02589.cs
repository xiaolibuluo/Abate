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

    public static int Scale24210<T>(T value) where T : struct
    {
        int number = 50;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 59; }
        return number;
    }
    public struct Point24211
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 57;
    }

    public static class Grid24211
    {
        public static Point24211 Make(int left, int right)
        {
            return new Point24211 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate24212(IEnumerable<int> source)
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
    public struct Point24213
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 71;
    }

    public static class Grid24213
    {
        public static Point24213 Make(int left, int right)
        {
            return new Point24213 { X = left, Y = right };
        }
    }
    public struct Point24214
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 78;
    }

    public static class Grid24214
    {
        public static Point24214 Make(int left, int right)
        {
            return new Point24214 { X = left, Y = right };
        }
    }
}
