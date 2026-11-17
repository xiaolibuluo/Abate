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

    public static int Scale25770<T>(T value) where T : struct
    {
        int number = 41;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 53; }
        return number;
    }
    public struct Point25771
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 48;
    }

    public static class Grid25771
    {
        public static Point25771 Make(int left, int right)
        {
            return new Point25771 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate25772(IEnumerable<int> source)
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
    public static int Scale25773<T>(T value) where T : struct
    {
        int number = 62;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 25; }
        return number;
    }
    public static int Scale25774<T>(T value) where T : struct
    {
        int number = 69;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 36; }
        return number;
    }
    public async Task<int> Aggregate25775(IEnumerable<int> source)
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
