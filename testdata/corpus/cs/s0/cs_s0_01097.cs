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

    public delegate void Notifier10290(int value);

    public sealed class Source10290
    {
        public event Notifier10290 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point10291
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 100;
    }

    public static class Grid10291
    {
        public static Point10291 Make(int left, int right)
        {
            return new Point10291 { X = left, Y = right };
        }
    }
}
