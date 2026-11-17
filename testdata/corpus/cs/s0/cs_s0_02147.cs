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

    public delegate void Notifier20060(int value);

    public sealed class Source20060
    {
        public event Notifier20060 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point20061
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 94;
    }

    public static class Grid20061
    {
        public static Point20061 Make(int left, int right)
        {
            return new Point20061 { X = left, Y = right };
        }
    }
}
