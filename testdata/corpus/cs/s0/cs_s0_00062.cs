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

    public static int Scale590<T>(T value) where T : struct
    {
        int number = 93;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 56; }
        return number;
    }
    public struct Point591
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 100;
    }

    public static class Grid591
    {
        public static Point591 Make(int left, int right)
        {
            return new Point591 { X = left, Y = right };
        }
    }
}
