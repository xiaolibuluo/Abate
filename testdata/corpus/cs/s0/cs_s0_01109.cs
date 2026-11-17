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

    public static int Scale10390<T>(T value) where T : struct
    {
        int number = 29;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 23; }
        return number;
    }
    public static int Scale10391<T>(T value) where T : struct
    {
        int number = 36;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 34; }
        return number;
    }
    public struct Point10392
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 43;
    }

    public static class Grid10392
    {
        public static Point10392 Make(int left, int right)
        {
            return new Point10392 { X = left, Y = right };
        }
    }
}
