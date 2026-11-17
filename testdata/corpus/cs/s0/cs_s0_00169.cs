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

    public delegate void Notifier1610(int value);

    public sealed class Source1610
    {
        public event Notifier1610 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier1611(int value);

    public sealed class Source1611
    {
        public event Notifier1611 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point1612
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 75;
    }

    public static class Grid1612
    {
        public static Point1612 Make(int left, int right)
        {
            return new Point1612 { X = left, Y = right };
        }
    }
}
