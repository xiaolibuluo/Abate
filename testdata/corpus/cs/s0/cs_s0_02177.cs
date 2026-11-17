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

    public static int Scale20320<T>(T value) where T : struct
    {
        int number = 37;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 20; }
        return number;
    }
    public struct Point20321
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 44;
    }

    public static class Grid20321
    {
        public static Point20321 Make(int left, int right)
        {
            return new Point20321 { X = left, Y = right };
        }
    }
    public static int Scale20322<T>(T value) where T : struct
    {
        int number = 51;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 42; }
        return number;
    }
    public static int Scale20323<T>(T value) where T : struct
    {
        int number = 58;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 53; }
        return number;
    }
    public struct Point20324
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 65;
    }

    public static class Grid20324
    {
        public static Point20324 Make(int left, int right)
        {
            return new Point20324 { X = left, Y = right };
        }
    }
}
