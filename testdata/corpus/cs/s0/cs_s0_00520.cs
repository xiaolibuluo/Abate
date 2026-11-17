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

    public struct Point4870
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 31;
    }

    public static class Grid4870
    {
        public static Point4870 Make(int left, int right)
        {
            return new Point4870 { X = left, Y = right };
        }
    }
}
