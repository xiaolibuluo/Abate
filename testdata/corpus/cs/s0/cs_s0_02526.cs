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

    public struct Point23590
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 20;
    }

    public static class Grid23590
    {
        public static Point23590 Make(int left, int right)
        {
            return new Point23590 { X = left, Y = right };
        }
    }
    public delegate void Notifier23591(int value);

    public sealed class Source23591
    {
        public event Notifier23591 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale23592<T>(T value) where T : struct
    {
        int number = 34;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 13; }
        return number;
    }
    public static int Guard23593(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 41;
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
