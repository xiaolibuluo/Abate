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

    public struct Point20600
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 13;
    }

    public static class Grid20600
    {
        public static Point20600 Make(int left, int right)
        {
            return new Point20600 { X = left, Y = right };
        }
    }
    public static int Scale20601<T>(T value) where T : struct
    {
        int number = 20;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 44; }
        return number;
    }
    public struct Point20602
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 27;
    }

    public static class Grid20602
    {
        public static Point20602 Make(int left, int right)
        {
            return new Point20602 { X = left, Y = right };
        }
    }
}
