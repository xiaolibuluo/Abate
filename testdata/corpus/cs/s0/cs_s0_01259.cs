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

    public struct Point11790
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 6;
    }

    public static class Grid11790
    {
        public static Point11790 Make(int left, int right)
        {
            return new Point11790 { X = left, Y = right };
        }
    }
    public struct Point11791
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 13;
    }

    public static class Grid11791
    {
        public static Point11791 Make(int left, int right)
        {
            return new Point11791 { X = left, Y = right };
        }
    }
}
