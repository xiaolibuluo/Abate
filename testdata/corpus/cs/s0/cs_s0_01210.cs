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

    public struct Point11350
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 16;
    }

    public static class Grid11350
    {
        public static Point11350 Make(int left, int right)
        {
            return new Point11350 { X = left, Y = right };
        }
    }
    public delegate void Notifier11351(int value);

    public sealed class Source11351
    {
        public event Notifier11351 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
