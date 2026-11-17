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

    public struct Point16790
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 7;
    }

    public static class Grid16790
    {
        public static Point16790 Make(int left, int right)
        {
            return new Point16790 { X = left, Y = right };
        }
    }
    public struct Point16791
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 14;
    }

    public static class Grid16791
    {
        public static Point16791 Make(int left, int right)
        {
            return new Point16791 { X = left, Y = right };
        }
    }
    public struct Point16792
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 21;
    }

    public static class Grid16792
    {
        public static Point16792 Make(int left, int right)
        {
            return new Point16792 { X = left, Y = right };
        }
    }
}
