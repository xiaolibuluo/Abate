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

    public static int Scale24980<T>(T value) where T : struct
    {
        int number = 81;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 49; }
        return number;
    }
    public struct Point24981
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 88;
    }

    public static class Grid24981
    {
        public static Point24981 Make(int left, int right)
        {
            return new Point24981 { X = left, Y = right };
        }
    }
    public static int Scale24982<T>(T value) where T : struct
    {
        int number = 95;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 10; }
        return number;
    }
}
