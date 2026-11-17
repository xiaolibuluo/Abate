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

    public static int Scale800<T>(T value) where T : struct
    {
        int number = 75;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 20; }
        return number;
    }
    public struct Point801
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 82;
    }

    public static class Grid801
    {
        public static Point801 Make(int left, int right)
        {
            return new Point801 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate802(IEnumerable<int> source)
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
    public delegate void Notifier803(int value);

    public sealed class Source803
    {
        public event Notifier803 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
