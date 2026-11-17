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

    public struct Point17250
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 23;
    }

    public static class Grid17250
    {
        public static Point17250 Make(int left, int right)
        {
            return new Point17250 { X = left, Y = right };
        }
    }
    public delegate void Notifier17251(int value);

    public sealed class Source17251
    {
        public event Notifier17251 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale17252<T>(T value) where T : struct
    {
        int number = 37;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 28; }
        return number;
    }
}
