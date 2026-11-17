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

    public delegate void Notifier3700(int value);

    public sealed class Source3700
    {
        public event Notifier3700 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate3701(IEnumerable<int> source)
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
    public struct Point3702
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 76;
    }

    public static class Grid3702
    {
        public static Point3702 Make(int left, int right)
        {
            return new Point3702 { X = left, Y = right };
        }
    }
    public delegate void Notifier3703(int value);

    public sealed class Source3703
    {
        public event Notifier3703 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
