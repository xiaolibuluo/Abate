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

    public delegate void Notifier18280(int value);

    public sealed class Source18280
    {
        public event Notifier18280 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier18281(int value);

    public sealed class Source18281
    {
        public event Notifier18281 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point18282
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 18;
    }

    public static class Grid18282
    {
        public static Point18282 Make(int left, int right)
        {
            return new Point18282 { X = left, Y = right };
        }
    }
    public struct Point18283
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 25;
    }

    public static class Grid18283
    {
        public static Point18283 Make(int left, int right)
        {
            return new Point18283 { X = left, Y = right };
        }
    }
}
