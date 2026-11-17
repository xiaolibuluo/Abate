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

    public static int Guard16800(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 20;
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
    public static int Scale16801<T>(T value) where T : struct
    {
        int number = 27;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 68; }
        return number;
    }
    public struct Point16802
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 34;
    }

    public static class Grid16802
    {
        public static Point16802 Make(int left, int right)
        {
            return new Point16802 { X = left, Y = right };
        }
    }
}
