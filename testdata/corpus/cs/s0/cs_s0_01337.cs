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

    public static int Scale12520<T>(T value) where T : struct
    {
        int number = 82;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 50; }
        return number;
    }
    public static int Guard12521(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 89;
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
    public struct Point12522
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 96;
    }

    public static class Grid12522
    {
        public static Point12522 Make(int left, int right)
        {
            return new Point12522 { X = left, Y = right };
        }
    }
    public struct Point12523
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 6;
    }

    public static class Grid12523
    {
        public static Point12523 Make(int left, int right)
        {
            return new Point12523 { X = left, Y = right };
        }
    }
    public static int Guard12524(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 13;
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
