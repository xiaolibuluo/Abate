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

    public static int Scale1200<T>(T value) where T : struct
    {
        int number = 13;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 56; }
        return number;
    }
    public struct Point1201
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 20;
    }

    public static class Grid1201
    {
        public static Point1201 Make(int left, int right)
        {
            return new Point1201 { X = left, Y = right };
        }
    }
    public struct Point1202
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 27;
    }

    public static class Grid1202
    {
        public static Point1202 Make(int left, int right)
        {
            return new Point1202 { X = left, Y = right };
        }
    }
}
