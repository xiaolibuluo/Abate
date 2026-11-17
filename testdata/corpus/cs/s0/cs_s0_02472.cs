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

    public struct Point23080
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 36;
    }

    public static class Grid23080
    {
        public static Point23080 Make(int left, int right)
        {
            return new Point23080 { X = left, Y = right };
        }
    }
    public struct Point23081
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 43;
    }

    public static class Grid23081
    {
        public static Point23081 Make(int left, int right)
        {
            return new Point23081 { X = left, Y = right };
        }
    }
    public delegate void Notifier23082(int value);

    public sealed class Source23082
    {
        public event Notifier23082 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
