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

    public delegate void Notifier15160(int value);

    public sealed class Source15160
    {
        public event Notifier15160 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point15161
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 29;
    }

    public static class Grid15161
    {
        public static Point15161 Make(int left, int right)
        {
            return new Point15161 { X = left, Y = right };
        }
    }
}
