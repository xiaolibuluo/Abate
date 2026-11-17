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

    public delegate void Notifier23460(int value);

    public sealed class Source23460
    {
        public event Notifier23460 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point23461
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 52;
    }

    public static class Grid23461
    {
        public static Point23461 Make(int left, int right)
        {
            return new Point23461 { X = left, Y = right };
        }
    }
}
