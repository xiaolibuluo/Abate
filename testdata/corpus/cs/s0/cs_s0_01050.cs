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

    public struct Point9830
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 77;
    }

    public static class Grid9830
    {
        public static Point9830 Make(int left, int right)
        {
            return new Point9830 { X = left, Y = right };
        }
    }
}
