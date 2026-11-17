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

    public struct Point10360
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 87;
    }

    public static class Grid10360
    {
        public static Point10360 Make(int left, int right)
        {
            return new Point10360 { X = left, Y = right };
        }
    }
    public static int Scale10361<T>(T value) where T : struct
    {
        int number = 94;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 13; }
        return number;
    }
}
