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

    public delegate void Notifier2820(int value);

    public sealed class Source2820
    {
        public event Notifier2820 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier2821(int value);

    public sealed class Source2821
    {
        public event Notifier2821 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard2822(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 96;
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
    public delegate void Notifier2823(int value);

    public sealed class Source2823
    {
        public event Notifier2823 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale2824<T>(T value) where T : struct
    {
        int number = 13;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 14; }
        return number;
    }
    public struct Point2825
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 20;
    }

    public static class Grid2825
    {
        public static Point2825 Make(int left, int right)
        {
            return new Point2825 { X = left, Y = right };
        }
    }
}
