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

    public struct Point9440
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 55;
    }

    public static class Grid9440
    {
        public static Point9440 Make(int left, int right)
        {
            return new Point9440 { X = left, Y = right };
        }
    }
    public delegate void Notifier9441(int value);

    public sealed class Source9441
    {
        public event Notifier9441 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point9442
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 69;
    }

    public static class Grid9442
    {
        public static Point9442 Make(int left, int right)
        {
            return new Point9442 { X = left, Y = right };
        }
    }
}
