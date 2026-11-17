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

    public struct Point26430
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 26;
    }

    public static class Grid26430
    {
        public static Point26430 Make(int left, int right)
        {
            return new Point26430 { X = left, Y = right };
        }
    }
    public struct Point26431
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 33;
    }

    public static class Grid26431
    {
        public static Point26431 Make(int left, int right)
        {
            return new Point26431 { X = left, Y = right };
        }
    }
}
