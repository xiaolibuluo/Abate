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

    public struct Point12280
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 61;
    }

    public static class Grid12280
    {
        public static Point12280 Make(int left, int right)
        {
            return new Point12280 { X = left, Y = right };
        }
    }
    public struct Point12281
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 68;
    }

    public static class Grid12281
    {
        public static Point12281 Make(int left, int right)
        {
            return new Point12281 { X = left, Y = right };
        }
    }
}
