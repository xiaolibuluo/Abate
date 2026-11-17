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

    public delegate void Notifier9010(int value);

    public sealed class Source9010
    {
        public event Notifier9010 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point9011
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 85;
    }

    public static class Grid9011
    {
        public static Point9011 Make(int left, int right)
        {
            return new Point9011 { X = left, Y = right };
        }
    }
    public delegate void Notifier9012(int value);

    public sealed class Source9012
    {
        public event Notifier9012 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale9013<T>(T value) where T : struct
    {
        int number = 99;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 66; }
        return number;
    }
}
