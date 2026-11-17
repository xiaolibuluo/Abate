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

    public struct Point19920
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 99;
    }

    public static class Grid19920
    {
        public static Point19920 Make(int left, int right)
        {
            return new Point19920 { X = left, Y = right };
        }
    }
    public static int Scale19921<T>(T value) where T : struct
    {
        int number = 9;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 56; }
        return number;
    }
    public static int Guard19922(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 16;
        }
        catch (ArgumentException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("guard");
        }
    }
    public struct Point19923
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 23;
    }

    public static class Grid19923
    {
        public static Point19923 Make(int left, int right)
        {
            return new Point19923 { X = left, Y = right };
        }
    }
    public static int Scale19924<T>(T value) where T : struct
    {
        int number = 30;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 28; }
        return number;
    }
    public static int Scale19925<T>(T value) where T : struct
    {
        int number = 37;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 39; }
        return number;
    }
}
