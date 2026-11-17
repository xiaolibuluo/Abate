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

    public struct Point26230
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 57;
    }

    public static class Grid26230
    {
        public static Point26230 Make(int left, int right)
        {
            return new Point26230 { X = left, Y = right };
        }
    }
    public struct Point26231
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 64;
    }

    public static class Grid26231
    {
        public static Point26231 Make(int left, int right)
        {
            return new Point26231 { X = left, Y = right };
        }
    }
    public struct Point26232
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 71;
    }

    public static class Grid26232
    {
        public static Point26232 Make(int left, int right)
        {
            return new Point26232 { X = left, Y = right };
        }
    }
    public static int Scale26233<T>(T value) where T : struct
    {
        int number = 78;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 42; }
        return number;
    }
    public async Task<int> Aggregate26234(IEnumerable<int> source)
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
