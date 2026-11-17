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

    public struct Point11320
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 74;
    }

    public static class Grid11320
    {
        public static Point11320 Make(int left, int right)
        {
            return new Point11320 { X = left, Y = right };
        }
    }
    public delegate void Notifier11321(int value);

    public sealed class Source11321
    {
        public event Notifier11321 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
