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

    public struct Point23390
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 51;
    }

    public static class Grid23390
    {
        public static Point23390 Make(int left, int right)
        {
            return new Point23390 { X = left, Y = right };
        }
    }
}
