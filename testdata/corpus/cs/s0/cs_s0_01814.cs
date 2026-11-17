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

    public delegate void Notifier16930(int value);

    public sealed class Source16930
    {
        public event Notifier16930 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate16931(IEnumerable<int> source)
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
    public struct Point16932
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 9;
    }

    public static class Grid16932
    {
        public static Point16932 Make(int left, int right)
        {
            return new Point16932 { X = left, Y = right };
        }
    }
    public struct Point16933
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 16;
    }

    public static class Grid16933
    {
        public static Point16933 Make(int left, int right)
        {
            return new Point16933 { X = left, Y = right };
        }
    }
    public static int Guard16934(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 23;
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
