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

    public struct Point23820
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 28;
    }

    public static class Grid23820
    {
        public static Point23820 Make(int left, int right)
        {
            return new Point23820 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate23821(IEnumerable<int> source)
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
    public delegate void Notifier23822(int value);

    public sealed class Source23822
    {
        public event Notifier23822 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier23823(int value);

    public sealed class Source23823
    {
        public event Notifier23823 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale23824<T>(T value) where T : struct
    {
        int number = 56;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 13; }
        return number;
    }
    public struct Point23825
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 63;
    }

    public static class Grid23825
    {
        public static Point23825 Make(int left, int right)
        {
            return new Point23825 { X = left, Y = right };
        }
    }
}
