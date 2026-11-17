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

    public static int Scale25150<T>(T value) where T : struct
    {
        int number = 11;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 46; }
        return number;
    }
    public static int Scale25151<T>(T value) where T : struct
    {
        int number = 18;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 57; }
        return number;
    }
    public static int Scale25152<T>(T value) where T : struct
    {
        int number = 25;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 68; }
        return number;
    }
    public async Task<int> Aggregate25153(IEnumerable<int> source)
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
    public struct Point25154
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 39;
    }

    public static class Grid25154
    {
        public static Point25154 Make(int left, int right)
        {
            return new Point25154 { X = left, Y = right };
        }
    }
}
