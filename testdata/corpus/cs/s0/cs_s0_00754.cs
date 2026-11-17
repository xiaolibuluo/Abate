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

    public static int Scale7070<T>(T value) where T : struct
    {
        int number = 78;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 17; }
        return number;
    }
    public static int Scale7071<T>(T value) where T : struct
    {
        int number = 85;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 28; }
        return number;
    }
    public static int Guard7072(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 92;
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
    public struct Point7073
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 99;
    }

    public static class Grid7073
    {
        public static Point7073 Make(int left, int right)
        {
            return new Point7073 { X = left, Y = right };
        }
    }
}
