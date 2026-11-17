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

    public async Task<int> Aggregate23280(IEnumerable<int> source)
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
    public static int Scale23281<T>(T value) where T : struct
    {
        int number = 12;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 29; }
        return number;
    }
    public delegate void Notifier23282(int value);

    public sealed class Source23282
    {
        public event Notifier23282 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point23283
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 26;
    }

    public static class Grid23283
    {
        public static Point23283 Make(int left, int right)
        {
            return new Point23283 { X = left, Y = right };
        }
    }
    public static int Guard23284(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 33;
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
}
