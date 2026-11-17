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

    public struct Point26040
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 101;
    }

    public static class Grid26040
    {
        public static Point26040 Make(int left, int right)
        {
            return new Point26040 { X = left, Y = right };
        }
    }
    public static int Scale26041<T>(T value) where T : struct
    {
        int number = 11;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 9; }
        return number;
    }
    public delegate void Notifier26042(int value);

    public sealed class Source26042
    {
        public event Notifier26042 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale26043<T>(T value) where T : struct
    {
        int number = 25;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 31; }
        return number;
    }
}
