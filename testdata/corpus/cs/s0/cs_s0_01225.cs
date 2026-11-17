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

    public delegate void Notifier11490(int value);

    public sealed class Source11490
    {
        public event Notifier11490 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point11491
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 11;
    }

    public static class Grid11491
    {
        public static Point11491 Make(int left, int right)
        {
            return new Point11491 { X = left, Y = right };
        }
    }
    public delegate void Notifier11492(int value);

    public sealed class Source11492
    {
        public event Notifier11492 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale11493<T>(T value) where T : struct
    {
        int number = 25;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 33; }
        return number;
    }
    public static int Scale11494<T>(T value) where T : struct
    {
        int number = 32;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 44; }
        return number;
    }
}
