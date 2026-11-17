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

    public static int Guard11480(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 88;
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
    public static int Scale11481<T>(T value) where T : struct
    {
        int number = 95;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 65; }
        return number;
    }
    public struct Point11482
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 5;
    }

    public static class Grid11482
    {
        public static Point11482 Make(int left, int right)
        {
            return new Point11482 { X = left, Y = right };
        }
    }
    public static int Guard11483(int value)
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
}
