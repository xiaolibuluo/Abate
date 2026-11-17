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

    public struct Point19300
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 69;
    }

    public static class Grid19300
    {
        public static Point19300 Make(int left, int right)
        {
            return new Point19300 { X = left, Y = right };
        }
    }
    public struct Point19301
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 76;
    }

    public static class Grid19301
    {
        public static Point19301 Make(int left, int right)
        {
            return new Point19301 { X = left, Y = right };
        }
    }
    public struct Point19302
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 83;
    }

    public static class Grid19302
    {
        public static Point19302 Make(int left, int right)
        {
            return new Point19302 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate19303(IEnumerable<int> source)
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
