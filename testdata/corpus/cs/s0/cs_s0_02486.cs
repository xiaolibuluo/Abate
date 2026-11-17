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

    public static int Scale23210<T>(T value) where T : struct
    {
        int number = 11;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 30; }
        return number;
    }
    public struct Point23211
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 18;
    }

    public static class Grid23211
    {
        public static Point23211 Make(int left, int right)
        {
            return new Point23211 { X = left, Y = right };
        }
    }
    public delegate void Notifier23212(int value);

    public sealed class Source23212
    {
        public event Notifier23212 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale23213<T>(T value) where T : struct
    {
        int number = 32;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 63; }
        return number;
    }
    public struct Point23214
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 39;
    }

    public static class Grid23214
    {
        public static Point23214 Make(int left, int right)
        {
            return new Point23214 { X = left, Y = right };
        }
    }
}
