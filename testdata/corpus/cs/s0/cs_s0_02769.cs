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

    public struct Point25900
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 16;
    }

    public static class Grid25900
    {
        public static Point25900 Make(int left, int right)
        {
            return new Point25900 { X = left, Y = right };
        }
    }
    public static int Scale25901<T>(T value) where T : struct
    {
        int number = 23;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 33; }
        return number;
    }
}
