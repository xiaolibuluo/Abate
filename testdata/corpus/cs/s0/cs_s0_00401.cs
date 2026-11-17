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

    public static int Scale3780<T>(T value) where T : struct
    {
        int number = 69;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 32; }
        return number;
    }
    public delegate void Notifier3781(int value);

    public sealed class Source3781
    {
        public event Notifier3781 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point3782
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 83;
    }

    public static class Grid3782
    {
        public static Point3782 Make(int left, int right)
        {
            return new Point3782 { X = left, Y = right };
        }
    }
    public struct Point3783
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 90;
    }

    public static class Grid3783
    {
        public static Point3783 Make(int left, int right)
        {
            return new Point3783 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate3784(IEnumerable<int> source)
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
}
