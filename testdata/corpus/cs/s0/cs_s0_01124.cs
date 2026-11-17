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

    public delegate void Notifier10540(int value);

    public sealed class Source10540
    {
        public event Notifier10540 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate10541(IEnumerable<int> source)
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
    public struct Point10542
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 44;
    }

    public static class Grid10542
    {
        public static Point10542 Make(int left, int right)
        {
            return new Point10542 { X = left, Y = right };
        }
    }
    public delegate void Notifier10543(int value);

    public sealed class Source10543
    {
        public event Notifier10543 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
