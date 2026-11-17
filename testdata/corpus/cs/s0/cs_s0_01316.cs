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

    public struct Point12320
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 16;
    }

    public static class Grid12320
    {
        public static Point12320 Make(int left, int right)
        {
            return new Point12320 { X = left, Y = right };
        }
    }
    public delegate void Notifier12321(int value);

    public sealed class Source12321
    {
        public event Notifier12321 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
