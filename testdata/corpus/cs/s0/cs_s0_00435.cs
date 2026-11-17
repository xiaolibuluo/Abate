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

    public struct Point4090
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 84;
    }

    public static class Grid4090
    {
        public static Point4090 Make(int left, int right)
        {
            return new Point4090 { X = left, Y = right };
        }
    }
    public delegate void Notifier4091(int value);

    public sealed class Source4091
    {
        public event Notifier4091 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point4092
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 98;
    }

    public static class Grid4092
    {
        public static Point4092 Make(int left, int right)
        {
            return new Point4092 { X = left, Y = right };
        }
    }
}
