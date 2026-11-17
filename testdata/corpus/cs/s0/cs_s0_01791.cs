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

    public struct Point16710
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 97;
    }

    public static class Grid16710
    {
        public static Point16710 Make(int left, int right)
        {
            return new Point16710 { X = left, Y = right };
        }
    }
    public struct Point16711
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 7;
    }

    public static class Grid16711
    {
        public static Point16711 Make(int left, int right)
        {
            return new Point16711 { X = left, Y = right };
        }
    }
}
