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

    public async Task<int> Aggregate21830(IEnumerable<int> source)
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
    public struct Point21831
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 67;
    }

    public static class Grid21831
    {
        public static Point21831 Make(int left, int right)
        {
            return new Point21831 { X = left, Y = right };
        }
    }
    public struct Point21832
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 74;
    }

    public static class Grid21832
    {
        public static Point21832 Make(int left, int right)
        {
            return new Point21832 { X = left, Y = right };
        }
    }
    public struct Point21833
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 81;
    }

    public static class Grid21833
    {
        public static Point21833 Make(int left, int right)
        {
            return new Point21833 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate21834(IEnumerable<int> source)
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
