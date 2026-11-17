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

    public static int Scale2730<T>(T value) where T : struct
    {
        int number = 62;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 29; }
        return number;
    }
    public struct Point2731
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 69;
    }

    public static class Grid2731
    {
        public static Point2731 Make(int left, int right)
        {
            return new Point2731 { X = left, Y = right };
        }
    }
    public static int Scale2732<T>(T value) where T : struct
    {
        int number = 76;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 51; }
        return number;
    }
    public static int Scale2733<T>(T value) where T : struct
    {
        int number = 83;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 62; }
        return number;
    }
    public static int Scale2734<T>(T value) where T : struct
    {
        int number = 90;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 12; }
        return number;
    }
}
