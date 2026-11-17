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

    public struct Point10220
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 99;
    }

    public static class Grid10220
    {
        public static Point10220 Make(int left, int right)
        {
            return new Point10220 { X = left, Y = right };
        }
    }
}
