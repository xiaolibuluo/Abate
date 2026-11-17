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

    public static int Scale5760<T>(T value) where T : struct
    {
        int number = 24;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 15; }
        return number;
    }
    public delegate void Notifier5761(int value);

    public sealed class Source5761
    {
        public event Notifier5761 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard5762(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 38;
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
    public async Task<int> Aggregate5763(IEnumerable<int> source)
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
    public struct Point5764
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 52;
    }

    public static class Grid5764
    {
        public static Point5764 Make(int left, int right)
        {
            return new Point5764 { X = left, Y = right };
        }
    }
}
