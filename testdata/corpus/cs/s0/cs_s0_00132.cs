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

    public delegate void Notifier1280(int value);

    public sealed class Source1280
    {
        public event Notifier1280 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point1281
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 27;
    }

    public static class Grid1281
    {
        public static Point1281 Make(int left, int right)
        {
            return new Point1281 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate1282(IEnumerable<int> source)
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
    public delegate void Notifier1283(int value);

    public sealed class Source1283
    {
        public event Notifier1283 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
