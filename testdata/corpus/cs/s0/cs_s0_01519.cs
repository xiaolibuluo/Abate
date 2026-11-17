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

    public delegate void Notifier14230(int value);

    public sealed class Source14230
    {
        public event Notifier14230 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point14231
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 81;
    }

    public static class Grid14231
    {
        public static Point14231 Make(int left, int right)
        {
            return new Point14231 { X = left, Y = right };
        }
    }
}
