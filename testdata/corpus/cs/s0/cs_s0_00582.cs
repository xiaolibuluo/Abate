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

    public struct Point5450
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 9;
    }

    public static class Grid5450
    {
        public static Point5450 Make(int left, int right)
        {
            return new Point5450 { X = left, Y = right };
        }
    }
}
