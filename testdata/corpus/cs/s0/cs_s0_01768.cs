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

    public async Task<int> Aggregate16480(IEnumerable<int> source)
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
    public struct Point16481
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 96;
    }

    public static class Grid16481
    {
        public static Point16481 Make(int left, int right)
        {
            return new Point16481 { X = left, Y = right };
        }
    }
    public struct Point16482
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 6;
    }

    public static class Grid16482
    {
        public static Point16482 Make(int left, int right)
        {
            return new Point16482 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate16483(IEnumerable<int> source)
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
    public async Task<int> Aggregate16484(IEnumerable<int> source)
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
