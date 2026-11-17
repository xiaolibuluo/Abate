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

    public delegate void Notifier19690(int value);

    public sealed class Source19690
    {
        public event Notifier19690 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point19691
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 98;
    }

    public static class Grid19691
    {
        public static Point19691 Make(int left, int right)
        {
            return new Point19691 { X = left, Y = right };
        }
    }
    public delegate void Notifier19692(int value);

    public sealed class Source19692
    {
        public event Notifier19692 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier19693(int value);

    public sealed class Source19693
    {
        public event Notifier19693 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
