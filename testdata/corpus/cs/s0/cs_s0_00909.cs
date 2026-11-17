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

    public delegate void Notifier8530(int value);

    public sealed class Source8530
    {
        public event Notifier8530 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point8531
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 43;
    }

    public static class Grid8531
    {
        public static Point8531 Make(int left, int right)
        {
            return new Point8531 { X = left, Y = right };
        }
    }
    public struct Point8532
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 50;
    }

    public static class Grid8532
    {
        public static Point8532 Make(int left, int right)
        {
            return new Point8532 { X = left, Y = right };
        }
    }
    public static int Scale8533<T>(T value) where T : struct
    {
        int number = 57;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 35; }
        return number;
    }
}
