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

    public delegate void Notifier9090(int value);

    public sealed class Source9090
    {
        public event Notifier9090 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point9091
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 92;
    }

    public static class Grid9091
    {
        public static Point9091 Make(int left, int right)
        {
            return new Point9091 { X = left, Y = right };
        }
    }
    public delegate void Notifier9092(int value);

    public sealed class Source9092
    {
        public event Notifier9092 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier9093(int value);

    public sealed class Source9093
    {
        public event Notifier9093 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point9094
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 16;
    }

    public static class Grid9094
    {
        public static Point9094 Make(int left, int right)
        {
            return new Point9094 { X = left, Y = right };
        }
    }
}
