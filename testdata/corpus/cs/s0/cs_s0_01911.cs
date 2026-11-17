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

    public static int Scale17840<T>(T value) where T : struct
    {
        int number = 14;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 53; }
        return number;
    }
    public delegate void Notifier17841(int value);

    public sealed class Source17841
    {
        public event Notifier17841 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point17842
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 28;
    }

    public static class Grid17842
    {
        public static Point17842 Make(int left, int right)
        {
            return new Point17842 { X = left, Y = right };
        }
    }
    public delegate void Notifier17843(int value);

    public sealed class Source17843
    {
        public event Notifier17843 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier17844(int value);

    public sealed class Source17844
    {
        public event Notifier17844 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
