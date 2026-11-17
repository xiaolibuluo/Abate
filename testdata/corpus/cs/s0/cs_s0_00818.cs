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

    public async Task<int> Aggregate7670(IEnumerable<int> source)
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
    public static int Scale7671<T>(T value) where T : struct
    {
        int number = 89;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 21; }
        return number;
    }
    public struct Point7672
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 96;
    }

    public static class Grid7672
    {
        public static Point7672 Make(int left, int right)
        {
            return new Point7672 { X = left, Y = right };
        }
    }
    public delegate void Notifier7673(int value);

    public sealed class Source7673
    {
        public event Notifier7673 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point7674
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 13;
    }

    public static class Grid7674
    {
        public static Point7674 Make(int left, int right)
        {
            return new Point7674 { X = left, Y = right };
        }
    }
}
