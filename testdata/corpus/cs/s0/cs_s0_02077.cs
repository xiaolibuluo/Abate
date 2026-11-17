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

    public struct Point19370
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 63;
    }

    public static class Grid19370
    {
        public static Point19370 Make(int left, int right)
        {
            return new Point19370 { X = left, Y = right };
        }
    }
    public static int Scale19371<T>(T value) where T : struct
    {
        int number = 70;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 37; }
        return number;
    }
    public delegate void Notifier19372(int value);

    public sealed class Source19372
    {
        public event Notifier19372 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier19373(int value);

    public sealed class Source19373
    {
        public event Notifier19373 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale19374<T>(T value) where T : struct
    {
        int number = 91;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 9; }
        return number;
    }
    public struct Point19375
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 98;
    }

    public static class Grid19375
    {
        public static Point19375 Make(int left, int right)
        {
            return new Point19375 { X = left, Y = right };
        }
    }
}
