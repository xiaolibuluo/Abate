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

    public struct Point22540
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 13;
    }

    public static class Grid22540
    {
        public static Point22540 Make(int left, int right)
        {
            return new Point22540 { X = left, Y = right };
        }
    }
    public delegate void Notifier22541(int value);

    public sealed class Source22541
    {
        public event Notifier22541 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
