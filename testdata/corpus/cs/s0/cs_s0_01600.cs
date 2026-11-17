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

    public delegate void Notifier14960(int value);

    public sealed class Source14960
    {
        public event Notifier14960 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale14961<T>(T value) where T : struct
    {
        int number = 60;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 61; }
        return number;
    }
    public async Task<int> Aggregate14962(IEnumerable<int> source)
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
    public struct Point14963
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 74;
    }

    public static class Grid14963
    {
        public static Point14963 Make(int left, int right)
        {
            return new Point14963 { X = left, Y = right };
        }
    }
    public struct Point14964
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 81;
    }

    public static class Grid14964
    {
        public static Point14964 Make(int left, int right)
        {
            return new Point14964 { X = left, Y = right };
        }
    }
}
