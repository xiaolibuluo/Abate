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

    public delegate void Notifier3340(int value);

    public sealed class Source3340
    {
        public event Notifier3340 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point3341
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 86;
    }

    public static class Grid3341
    {
        public static Point3341 Make(int left, int right)
        {
            return new Point3341 { X = left, Y = right };
        }
    }
    public struct Point3342
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 93;
    }

    public static class Grid3342
    {
        public static Point3342 Make(int left, int right)
        {
            return new Point3342 { X = left, Y = right };
        }
    }
    public struct Point3343
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 100;
    }

    public static class Grid3343
    {
        public static Point3343 Make(int left, int right)
        {
            return new Point3343 { X = left, Y = right };
        }
    }
    public delegate void Notifier3344(int value);

    public sealed class Source3344
    {
        public event Notifier3344 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier3345(int value);

    public sealed class Source3345
    {
        public event Notifier3345 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
