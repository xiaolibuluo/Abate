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

    public async Task<int> Aggregate3210(IEnumerable<int> source)
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
    public struct Point3211
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 14;
    }

    public static class Grid3211
    {
        public static Point3211 Make(int left, int right)
        {
            return new Point3211 { X = left, Y = right };
        }
    }
    public delegate void Notifier3212(int value);

    public sealed class Source3212
    {
        public event Notifier3212 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard3213(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 28;
        }
        catch (ArgumentException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("guard");
        }
    }
}
