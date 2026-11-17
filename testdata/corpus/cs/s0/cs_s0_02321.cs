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

    public delegate void Notifier21650(int value);

    public sealed class Source21650
    {
        public event Notifier21650 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale21651<T>(T value) where T : struct
    {
        int number = 27;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 47; }
        return number;
    }
    public delegate void Notifier21652(int value);

    public sealed class Source21652
    {
        public event Notifier21652 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point21653
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 41;
    }

    public static class Grid21653
    {
        public static Point21653 Make(int left, int right)
        {
            return new Point21653 { X = left, Y = right };
        }
    }
}
