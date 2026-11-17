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

    public static int Scale4290<T>(T value) where T : struct
    {
        int number = 53;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 23; }
        return number;
    }
    public delegate void Notifier4291(int value);

    public sealed class Source4291
    {
        public event Notifier4291 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier4292(int value);

    public sealed class Source4292
    {
        public event Notifier4292 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale4293<T>(T value) where T : struct
    {
        int number = 74;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 56; }
        return number;
    }
    public static int Guard4294(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 81;
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
    public struct Point4295
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 88;
    }

    public static class Grid4295
    {
        public static Point4295 Make(int left, int right)
        {
            return new Point4295 { X = left, Y = right };
        }
    }
}
