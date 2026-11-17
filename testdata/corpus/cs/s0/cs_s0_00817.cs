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

    public static int Guard7660(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 69;
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
    public static int Scale7661<T>(T value) where T : struct
    {
        int number = 76;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 14; }
        return number;
    }
    public struct Point7662
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 83;
    }

    public static class Grid7662
    {
        public static Point7662 Make(int left, int right)
        {
            return new Point7662 { X = left, Y = right };
        }
    }
    public struct Point7663
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 90;
    }

    public static class Grid7663
    {
        public static Point7663 Make(int left, int right)
        {
            return new Point7663 { X = left, Y = right };
        }
    }
    public static int Scale7664<T>(T value) where T : struct
    {
        int number = 97;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 47; }
        return number;
    }
}
