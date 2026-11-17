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

    public struct Point24180
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 11;
    }

    public static class Grid24180
    {
        public static Point24180 Make(int left, int right)
        {
            return new Point24180 { X = left, Y = right };
        }
    }
    public struct Point24181
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 18;
    }

    public static class Grid24181
    {
        public static Point24181 Make(int left, int right)
        {
            return new Point24181 { X = left, Y = right };
        }
    }
    public struct Point24182
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 25;
    }

    public static class Grid24182
    {
        public static Point24182 Make(int left, int right)
        {
            return new Point24182 { X = left, Y = right };
        }
    }
    public struct Point24183
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 32;
    }

    public static class Grid24183
    {
        public static Point24183 Make(int left, int right)
        {
            return new Point24183 { X = left, Y = right };
        }
    }
}
