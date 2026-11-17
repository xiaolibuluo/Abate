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

    public struct Point26090
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 69;
    }

    public static class Grid26090
    {
        public static Point26090 Make(int left, int right)
        {
            return new Point26090 { X = left, Y = right };
        }
    }
    public static int Guard26091(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 76;
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
    public delegate void Notifier26092(int value);

    public sealed class Source26092
    {
        public event Notifier26092 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate26093(IEnumerable<int> source)
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
