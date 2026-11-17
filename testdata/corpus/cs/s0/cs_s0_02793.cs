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

    public struct Point26120
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 11;
    }

    public static class Grid26120
    {
        public static Point26120 Make(int left, int right)
        {
            return new Point26120 { X = left, Y = right };
        }
    }
    public static int Scale26121<T>(T value) where T : struct
    {
        int number = 18;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 65; }
        return number;
    }
    public async Task<int> Aggregate26122(IEnumerable<int> source)
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
    public delegate void Notifier26123(int value);

    public sealed class Source26123
    {
        public event Notifier26123 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale26124<T>(T value) where T : struct
    {
        int number = 39;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 37; }
        return number;
    }
}
