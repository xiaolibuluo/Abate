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

    public struct Point21900
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 54;
    }

    public static class Grid21900
    {
        public static Point21900 Make(int left, int right)
        {
            return new Point21900 { X = left, Y = right };
        }
    }
}
