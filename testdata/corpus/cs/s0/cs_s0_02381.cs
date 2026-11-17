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

    public static int Guard22230(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 95;
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
    public struct Point22231
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 5;
    }

    public static class Grid22231
    {
        public static Point22231 Make(int left, int right)
        {
            return new Point22231 { X = left, Y = right };
        }
    }
    public static int Guard22232(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 12;
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
    public static int Scale22233<T>(T value) where T : struct
    {
        int number = 19;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 48; }
        return number;
    }
}
