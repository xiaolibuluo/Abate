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

    public static int Guard20180(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 49;
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
    public static int Scale20181<T>(T value) where T : struct
    {
        int number = 56;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 55; }
        return number;
    }
    public struct Point20182
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 63;
    }

    public static class Grid20182
    {
        public static Point20182 Make(int left, int right)
        {
            return new Point20182 { X = left, Y = right };
        }
    }
}
