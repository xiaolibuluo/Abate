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

    public delegate void Notifier8030(int value);

    public sealed class Source8030
    {
        public event Notifier8030 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier8031(int value);

    public sealed class Source8031
    {
        public event Notifier8031 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate8032(IEnumerable<int> source)
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
    public struct Point8033
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 86;
    }

    public static class Grid8033
    {
        public static Point8033 Make(int left, int right)
        {
            return new Point8033 { X = left, Y = right };
        }
    }
}
