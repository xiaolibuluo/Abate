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

    public static int Guard22330(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 31;
        }
        catch (ArgumentException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("guard");
        }
    }
    public async Task<int> Aggregate22331(IEnumerable<int> source)
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
    public struct Point22332
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 45;
    }

    public static class Grid22332
    {
        public static Point22332 Make(int left, int right)
        {
            return new Point22332 { X = left, Y = right };
        }
    }
    public static int Scale22333<T>(T value) where T : struct
    {
        int number = 52;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 57; }
        return number;
    }
    public struct Point22334
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 59;
    }

    public static class Grid22334
    {
        public static Point22334 Make(int left, int right)
        {
            return new Point22334 { X = left, Y = right };
        }
    }
}
