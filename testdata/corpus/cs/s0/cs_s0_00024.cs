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

    public static int Scale230<T>(T value) where T : struct
    {
        int number = 13;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 48; }
        return number;
    }
    public delegate void Notifier231(int value);

    public sealed class Source231
    {
        public event Notifier231 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale232<T>(T value) where T : struct
    {
        int number = 27;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 9; }
        return number;
    }
    public struct Point233
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 34;
    }

    public static class Grid233
    {
        public static Point233 Make(int left, int right)
        {
            return new Point233 { X = left, Y = right };
        }
    }
}
