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

    public static int Scale7150<T>(T value) where T : struct
    {
        int number = 85;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 12; }
        return number;
    }
    public static int Scale7151<T>(T value) where T : struct
    {
        int number = 92;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 23; }
        return number;
    }
    public struct Point7152
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 99;
    }

    public static class Grid7152
    {
        public static Point7152 Make(int left, int right)
        {
            return new Point7152 { X = left, Y = right };
        }
    }
    public struct Point7153
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 9;
    }

    public static class Grid7153
    {
        public static Point7153 Make(int left, int right)
        {
            return new Point7153 { X = left, Y = right };
        }
    }
    public delegate void Notifier7154(int value);

    public sealed class Source7154
    {
        public event Notifier7154 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
