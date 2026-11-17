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

    public delegate void Notifier24190(int value);

    public sealed class Source24190
    {
        public event Notifier24190 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier24191(int value);

    public sealed class Source24191
    {
        public event Notifier24191 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point24192
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 38;
    }

    public static class Grid24192
    {
        public static Point24192 Make(int left, int right)
        {
            return new Point24192 { X = left, Y = right };
        }
    }
    public static int Scale24193<T>(T value) where T : struct
    {
        int number = 45;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 17; }
        return number;
    }
    public static int Guard24194(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 52;
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
