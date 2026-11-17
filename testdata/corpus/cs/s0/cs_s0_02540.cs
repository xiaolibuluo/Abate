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

    public async Task<int> Aggregate23730(IEnumerable<int> source)
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
    public async Task<int> Aggregate23731(IEnumerable<int> source)
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
    public delegate void Notifier23732(int value);

    public sealed class Source23732
    {
        public event Notifier23732 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point23733
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 29;
    }

    public static class Grid23733
    {
        public static Point23733 Make(int left, int right)
        {
            return new Point23733 { X = left, Y = right };
        }
    }
}
