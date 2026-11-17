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

    public delegate void Notifier7180(int value);

    public sealed class Source7180
    {
        public event Notifier7180 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate7181(IEnumerable<int> source)
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
    public struct Point7182
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 41;
    }

    public static class Grid7182
    {
        public static Point7182 Make(int left, int right)
        {
            return new Point7182 { X = left, Y = right };
        }
    }
    public delegate void Notifier7183(int value);

    public sealed class Source7183
    {
        public event Notifier7183 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale7184<T>(T value) where T : struct
    {
        int number = 55;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 16; }
        return number;
    }
}
