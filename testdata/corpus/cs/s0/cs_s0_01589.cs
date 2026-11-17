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

    public delegate void Notifier14860(int value);

    public sealed class Source14860
    {
        public event Notifier14860 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale14861<T>(T value) where T : struct
    {
        int number = 27;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 52; }
        return number;
    }
    public struct Point14862
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 34;
    }

    public static class Grid14862
    {
        public static Point14862 Make(int left, int right)
        {
            return new Point14862 { X = left, Y = right };
        }
    }
    public struct Point14863
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 41;
    }

    public static class Grid14863
    {
        public static Point14863 Make(int left, int right)
        {
            return new Point14863 { X = left, Y = right };
        }
    }
}
