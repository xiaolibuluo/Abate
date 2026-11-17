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

    public struct Point13380
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 36;
    }

    public static class Grid13380
    {
        public static Point13380 Make(int left, int right)
        {
            return new Point13380 { X = left, Y = right };
        }
    }
    public delegate void Notifier13381(int value);

    public sealed class Source13381
    {
        public event Notifier13381 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point13382
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 50;
    }

    public static class Grid13382
    {
        public static Point13382 Make(int left, int right)
        {
            return new Point13382 { X = left, Y = right };
        }
    }
}
