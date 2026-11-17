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

    public delegate void Notifier21130(int value);

    public sealed class Source21130
    {
        public event Notifier21130 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale21131<T>(T value) where T : struct
    {
        int number = 30;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 49; }
        return number;
    }
    public struct Point21132
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 37;
    }

    public static class Grid21132
    {
        public static Point21132 Make(int left, int right)
        {
            return new Point21132 { X = left, Y = right };
        }
    }
}
