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

    public delegate void Notifier7430(int value);

    public sealed class Source7430
    {
        public event Notifier7430 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point7431
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 68;
    }

    public static class Grid7431
    {
        public static Point7431 Make(int left, int right)
        {
            return new Point7431 { X = left, Y = right };
        }
    }
}
