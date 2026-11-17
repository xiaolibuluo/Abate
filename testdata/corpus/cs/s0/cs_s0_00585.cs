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

    public delegate void Notifier5480(int value);

    public sealed class Source5480
    {
        public event Notifier5480 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point5481
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 55;
    }

    public static class Grid5481
    {
        public static Point5481 Make(int left, int right)
        {
            return new Point5481 { X = left, Y = right };
        }
    }
}
