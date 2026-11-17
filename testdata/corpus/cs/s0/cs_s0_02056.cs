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

    public static int Scale19190<T>(T value) where T : struct
    {
        int number = 23;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 22; }
        return number;
    }
    public struct Point19191
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 30;
    }

    public static class Grid19191
    {
        public static Point19191 Make(int left, int right)
        {
            return new Point19191 { X = left, Y = right };
        }
    }
    public static int Scale19192<T>(T value) where T : struct
    {
        int number = 37;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 44; }
        return number;
    }
}
