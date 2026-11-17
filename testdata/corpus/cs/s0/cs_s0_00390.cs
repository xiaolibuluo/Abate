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

    public static int Scale3690<T>(T value) where T : struct
    {
        int number = 49;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 30; }
        return number;
    }
    public static int Scale3691<T>(T value) where T : struct
    {
        int number = 56;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 41; }
        return number;
    }
    public struct Point3692
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 63;
    }

    public static class Grid3692
    {
        public static Point3692 Make(int left, int right)
        {
            return new Point3692 { X = left, Y = right };
        }
    }
    public static int Scale3693<T>(T value) where T : struct
    {
        int number = 70;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 63; }
        return number;
    }
}
