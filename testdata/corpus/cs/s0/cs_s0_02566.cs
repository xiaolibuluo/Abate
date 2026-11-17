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

    public struct Point23990
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 55;
    }

    public static class Grid23990
    {
        public static Point23990 Make(int left, int right)
        {
            return new Point23990 { X = left, Y = right };
        }
    }
    public static int Scale23991<T>(T value) where T : struct
    {
        int number = 62;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 38; }
        return number;
    }
}
