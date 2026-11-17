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

    public static int Scale20800<T>(T value) where T : struct
    {
        int number = 79;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 51; }
        return number;
    }
    public static int Scale20801<T>(T value) where T : struct
    {
        int number = 86;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 62; }
        return number;
    }
    public static int Guard20802(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 93;
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
    public delegate void Notifier20803(int value);

    public sealed class Source20803
    {
        public event Notifier20803 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point20804
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 10;
    }

    public static class Grid20804
    {
        public static Point20804 Make(int left, int right)
        {
            return new Point20804 { X = left, Y = right };
        }
    }
}
