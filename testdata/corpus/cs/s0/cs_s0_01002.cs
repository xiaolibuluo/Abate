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

    public static int Scale9400<T>(T value) where T : struct
    {
        int number = 100;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 62; }
        return number;
    }
    public struct Point9401
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 10;
    }

    public static class Grid9401
    {
        public static Point9401 Make(int left, int right)
        {
            return new Point9401 { X = left, Y = right };
        }
    }
}
