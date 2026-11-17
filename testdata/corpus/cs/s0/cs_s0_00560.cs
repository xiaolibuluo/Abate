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

    public static int Scale5250<T>(T value) where T : struct
    {
        int number = 40;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 24; }
        return number;
    }
    public static int Guard5251(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 47;
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
    public delegate void Notifier5252(int value);

    public sealed class Source5252
    {
        public event Notifier5252 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier5253(int value);

    public sealed class Source5253
    {
        public event Notifier5253 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point5254
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 68;
    }

    public static class Grid5254
    {
        public static Point5254 Make(int left, int right)
        {
            return new Point5254 { X = left, Y = right };
        }
    }
    public delegate void Notifier5255(int value);

    public sealed class Source5255
    {
        public event Notifier5255 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
