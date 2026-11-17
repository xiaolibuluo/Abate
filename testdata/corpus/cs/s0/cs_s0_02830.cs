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

    public static int Scale26470<T>(T value) where T : struct
    {
        int number = 78;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 55; }
        return number;
    }
    public delegate void Notifier26471(int value);

    public sealed class Source26471
    {
        public event Notifier26471 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point26472
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 92;
    }

    public static class Grid26472
    {
        public static Point26472 Make(int left, int right)
        {
            return new Point26472 { X = left, Y = right };
        }
    }
}
