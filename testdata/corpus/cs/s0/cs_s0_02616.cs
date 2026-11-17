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

    public struct Point24470
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 97;
    }

    public static class Grid24470
    {
        public static Point24470 Make(int left, int right)
        {
            return new Point24470 { X = left, Y = right };
        }
    }
    public static int Scale24471<T>(T value) where T : struct
    {
        int number = 7;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 69; }
        return number;
    }
}
