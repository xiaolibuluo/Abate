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

    public struct Point5490
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 61;
    }

    public static class Grid5490
    {
        public static Point5490 Make(int left, int right)
        {
            return new Point5490 { X = left, Y = right };
        }
    }
    public delegate void Notifier5491(int value);

    public sealed class Source5491
    {
        public event Notifier5491 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
