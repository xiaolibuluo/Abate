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

    public delegate void Notifier22990(int value);

    public sealed class Source22990
    {
        public event Notifier22990 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point22991
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 23;
    }

    public static class Grid22991
    {
        public static Point22991 Make(int left, int right)
        {
            return new Point22991 { X = left, Y = right };
        }
    }
}
