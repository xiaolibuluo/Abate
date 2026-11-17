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

    public struct Point9040
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 20;
    }

    public static class Grid9040
    {
        public static Point9040 Make(int left, int right)
        {
            return new Point9040 { X = left, Y = right };
        }
    }
    public struct Point9041
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 27;
    }

    public static class Grid9041
    {
        public static Point9041 Make(int left, int right)
        {
            return new Point9041 { X = left, Y = right };
        }
    }
    public delegate void Notifier9042(int value);

    public sealed class Source9042
    {
        public event Notifier9042 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point9043
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 41;
    }

    public static class Grid9043
    {
        public static Point9043 Make(int left, int right)
        {
            return new Point9043 { X = left, Y = right };
        }
    }
}
