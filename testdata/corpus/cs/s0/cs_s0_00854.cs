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

    public delegate void Notifier8010(int value);

    public sealed class Source8010
    {
        public event Notifier8010 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier8011(int value);

    public sealed class Source8011
    {
        public event Notifier8011 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point8012
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 53;
    }

    public static class Grid8012
    {
        public static Point8012 Make(int left, int right)
        {
            return new Point8012 { X = left, Y = right };
        }
    }
    public static int Scale8013<T>(T value) where T : struct
    {
        int number = 60;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 37; }
        return number;
    }
}
