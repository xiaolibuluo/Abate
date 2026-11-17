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

    public static int Scale24200<T>(T value) where T : struct
    {
        int number = 37;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 52; }
        return number;
    }
    public struct Point24201
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 44;
    }

    public static class Grid24201
    {
        public static Point24201 Make(int left, int right)
        {
            return new Point24201 { X = left, Y = right };
        }
    }
    public static int Guard24202(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 51;
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
    public struct Point24203
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 58;
    }

    public static class Grid24203
    {
        public static Point24203 Make(int left, int right)
        {
            return new Point24203 { X = left, Y = right };
        }
    }
    public static int Scale24204<T>(T value) where T : struct
    {
        int number = 65;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 35; }
        return number;
    }
}
