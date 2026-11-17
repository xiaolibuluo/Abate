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

    public struct Point6560
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 94;
    }

    public static class Grid6560
    {
        public static Point6560 Make(int left, int right)
        {
            return new Point6560 { X = left, Y = right };
        }
    }
    public static int Scale6561<T>(T value) where T : struct
    {
        int number = 101;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 37; }
        return number;
    }
    public delegate void Notifier6562(int value);

    public sealed class Source6562
    {
        public event Notifier6562 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point6563
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 18;
    }

    public static class Grid6563
    {
        public static Point6563 Make(int left, int right)
        {
            return new Point6563 { X = left, Y = right };
        }
    }
}
