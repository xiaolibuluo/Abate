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

    public static int Scale5620<T>(T value) where T : struct
    {
        int number = 36;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 39; }
        return number;
    }
    public static int Scale5621<T>(T value) where T : struct
    {
        int number = 43;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 50; }
        return number;
    }
    public struct Point5622
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 50;
    }

    public static class Grid5622
    {
        public static Point5622 Make(int left, int right)
        {
            return new Point5622 { X = left, Y = right };
        }
    }
    public static int Scale5623<T>(T value) where T : struct
    {
        int number = 57;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 11; }
        return number;
    }
}
