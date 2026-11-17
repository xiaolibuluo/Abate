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

    public struct Point17640
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 45;
    }

    public static class Grid17640
    {
        public static Point17640 Make(int left, int right)
        {
            return new Point17640 { X = left, Y = right };
        }
    }
    public delegate void Notifier17641(int value);

    public sealed class Source17641
    {
        public event Notifier17641 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
