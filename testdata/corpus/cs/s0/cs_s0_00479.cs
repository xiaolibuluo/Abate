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

    public delegate void Notifier4490(int value);

    public sealed class Source4490
    {
        public event Notifier4490 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point4491
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 29;
    }

    public static class Grid4491
    {
        public static Point4491 Make(int left, int right)
        {
            return new Point4491 { X = left, Y = right };
        }
    }
}
