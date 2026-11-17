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

    public delegate void Notifier9840(int value);

    public sealed class Source9840
    {
        public event Notifier9840 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point9841
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 97;
    }

    public static class Grid9841
    {
        public static Point9841 Make(int left, int right)
        {
            return new Point9841 { X = left, Y = right };
        }
    }
}
