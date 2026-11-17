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

    public async Task<int> Aggregate950(IEnumerable<int> source)
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
    public delegate void Notifier951(int value);

    public sealed class Source951
    {
        public event Notifier951 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale952<T>(T value) where T : struct
    {
        int number = 90;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 25; }
        return number;
    }
    public async Task<int> Aggregate953(IEnumerable<int> source)
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
    public struct Point954
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 7;
    }

    public static class Grid954
    {
        public static Point954 Make(int left, int right)
        {
            return new Point954 { X = left, Y = right };
        }
    }
    public static int Guard955(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 14;
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
