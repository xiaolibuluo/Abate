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

    public struct Point6060
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 26;
    }

    public static class Grid6060
    {
        public static Point6060 Make(int left, int right)
        {
            return new Point6060 { X = left, Y = right };
        }
    }
    public delegate void Notifier6061(int value);

    public sealed class Source6061
    {
        public event Notifier6061 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier6062(int value);

    public sealed class Source6062
    {
        public event Notifier6062 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale6063<T>(T value) where T : struct
    {
        int number = 47;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 14; }
        return number;
    }
}
