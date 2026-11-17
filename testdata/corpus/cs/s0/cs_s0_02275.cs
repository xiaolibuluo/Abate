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

    public delegate void Notifier21250(int value);

    public sealed class Source21250
    {
        public event Notifier21250 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate21251(IEnumerable<int> source)
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
    public delegate void Notifier21252(int value);

    public sealed class Source21252
    {
        public event Notifier21252 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point21253
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 6;
    }

    public static class Grid21253
    {
        public static Point21253 Make(int left, int right)
        {
            return new Point21253 { X = left, Y = right };
        }
    }
    public delegate void Notifier21254(int value);

    public sealed class Source21254
    {
        public event Notifier21254 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
