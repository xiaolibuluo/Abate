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

    public static int Scale21980<T>(T value) where T : struct
    {
        int number = 61;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 23; }
        return number;
    }
    public struct Point21981
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 68;
    }

    public static class Grid21981
    {
        public static Point21981 Make(int left, int right)
        {
            return new Point21981 { X = left, Y = right };
        }
    }
}
