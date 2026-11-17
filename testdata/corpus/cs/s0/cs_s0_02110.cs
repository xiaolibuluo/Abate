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

    public static int Scale19700<T>(T value) where T : struct
    {
        int number = 7;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 13; }
        return number;
    }
    public struct Point19701
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 14;
    }

    public static class Grid19701
    {
        public static Point19701 Make(int left, int right)
        {
            return new Point19701 { X = left, Y = right };
        }
    }
    public struct Point19702
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 21;
    }

    public static class Grid19702
    {
        public static Point19702 Make(int left, int right)
        {
            return new Point19702 { X = left, Y = right };
        }
    }
}
