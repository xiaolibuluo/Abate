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

    public static int Scale21350<T>(T value) where T : struct
    {
        int number = 18;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 9; }
        return number;
    }
    public struct Point21351
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 25;
    }

    public static class Grid21351
    {
        public static Point21351 Make(int left, int right)
        {
            return new Point21351 { X = left, Y = right };
        }
    }
    public struct Point21352
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 32;
    }

    public static class Grid21352
    {
        public static Point21352 Make(int left, int right)
        {
            return new Point21352 { X = left, Y = right };
        }
    }
    public struct Point21353
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 39;
    }

    public static class Grid21353
    {
        public static Point21353 Make(int left, int right)
        {
            return new Point21353 { X = left, Y = right };
        }
    }
    public static int Guard21354(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 46;
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
}
