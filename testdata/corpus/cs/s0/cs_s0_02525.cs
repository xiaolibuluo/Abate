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

    public static int Scale23580<T>(T value) where T : struct
    {
        int number = 7;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 45; }
        return number;
    }
    public static int Guard23581(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 14;
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
    public struct Point23582
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 21;
    }

    public static class Grid23582
    {
        public static Point23582 Make(int left, int right)
        {
            return new Point23582 { X = left, Y = right };
        }
    }
}
