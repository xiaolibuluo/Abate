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

    public async Task<int> Aggregate22660(IEnumerable<int> source)
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
    public static int Scale22661<T>(T value) where T : struct
    {
        int number = 79;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 22; }
        return number;
    }
    public struct Point22662
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 86;
    }

    public static class Grid22662
    {
        public static Point22662 Make(int left, int right)
        {
            return new Point22662 { X = left, Y = right };
        }
    }
    public delegate void Notifier22663(int value);

    public sealed class Source22663
    {
        public event Notifier22663 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
