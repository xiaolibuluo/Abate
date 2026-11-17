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

    public static int Scale25080<T>(T value) where T : struct
    {
        int number = 17;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 58; }
        return number;
    }
    public static int Guard25081(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 24;
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
    public struct Point25082
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 31;
    }

    public static class Grid25082
    {
        public static Point25082 Make(int left, int right)
        {
            return new Point25082 { X = left, Y = right };
        }
    }
    public static int Guard25083(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 38;
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
