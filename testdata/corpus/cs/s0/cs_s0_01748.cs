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

    public delegate void Notifier16290(int value);

    public sealed class Source16290
    {
        public event Notifier16290 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point16291
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 43;
    }

    public static class Grid16291
    {
        public static Point16291 Make(int left, int right)
        {
            return new Point16291 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate16292(IEnumerable<int> source)
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
