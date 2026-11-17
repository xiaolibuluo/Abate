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

    public delegate void Notifier20730(int value);

    public sealed class Source20730
    {
        public event Notifier20730 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate20731(IEnumerable<int> source)
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
    public delegate void Notifier20732(int value);

    public sealed class Source20732
    {
        public event Notifier20732 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point20733
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 9;
    }

    public static class Grid20733
    {
        public static Point20733 Make(int left, int right)
        {
            return new Point20733 { X = left, Y = right };
        }
    }
}
