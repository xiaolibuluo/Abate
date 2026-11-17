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

    public delegate void Notifier8690(int value);

    public sealed class Source8690
    {
        public event Notifier8690 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale8691<T>(T value) where T : struct
    {
        int number = 57;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 64; }
        return number;
    }
    public delegate void Notifier8692(int value);

    public sealed class Source8692
    {
        public event Notifier8692 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point8693
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 71;
    }

    public static class Grid8693
    {
        public static Point8693 Make(int left, int right)
        {
            return new Point8693 { X = left, Y = right };
        }
    }
    public struct Point8694
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 78;
    }

    public static class Grid8694
    {
        public static Point8694 Make(int left, int right)
        {
            return new Point8694 { X = left, Y = right };
        }
    }
    public struct Point8695
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 85;
    }

    public static class Grid8695
    {
        public static Point8695 Make(int left, int right)
        {
            return new Point8695 { X = left, Y = right };
        }
    }
}
