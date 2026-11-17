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

    public struct Point17380
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 95;
    }

    public static class Grid17380
    {
        public static Point17380 Make(int left, int right)
        {
            return new Point17380 { X = left, Y = right };
        }
    }
    public struct Point17381
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 5;
    }

    public static class Grid17381
    {
        public static Point17381 Make(int left, int right)
        {
            return new Point17381 { X = left, Y = right };
        }
    }
    public static int Scale17382<T>(T value) where T : struct
    {
        int number = 12;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 58; }
        return number;
    }
    public static int Scale17383<T>(T value) where T : struct
    {
        int number = 19;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 69; }
        return number;
    }
    public struct Point17384
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 26;
    }

    public static class Grid17384
    {
        public static Point17384 Make(int left, int right)
        {
            return new Point17384 { X = left, Y = right };
        }
    }
}
