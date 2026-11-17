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

    public delegate void Notifier25580(int value);

    public sealed class Source25580
    {
        public event Notifier25580 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale25581<T>(T value) where T : struct
    {
        int number = 92;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 53; }
        return number;
    }
    public async Task<int> Aggregate25582(IEnumerable<int> source)
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
    public struct Point25583
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 9;
    }

    public static class Grid25583
    {
        public static Point25583 Make(int left, int right)
        {
            return new Point25583 { X = left, Y = right };
        }
    }
}
