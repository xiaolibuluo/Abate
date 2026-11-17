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

    public struct Point9430
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 42;
    }

    public static class Grid9430
    {
        public static Point9430 Make(int left, int right)
        {
            return new Point9430 { X = left, Y = right };
        }
    }
    public struct Point9431
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 49;
    }

    public static class Grid9431
    {
        public static Point9431 Make(int left, int right)
        {
            return new Point9431 { X = left, Y = right };
        }
    }
}
