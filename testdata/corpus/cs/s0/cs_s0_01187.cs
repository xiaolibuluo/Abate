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

    public struct Point11120
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 8;
    }

    public static class Grid11120
    {
        public static Point11120 Make(int left, int right)
        {
            return new Point11120 { X = left, Y = right };
        }
    }
    public static int Scale11121<T>(T value) where T : struct
    {
        int number = 15;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 57; }
        return number;
    }
    public interface IHandler11122
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope11122(string Label, int Payload);

    public sealed class Handler11122 : IHandler11122
    {
        public void Handle(int value) { }
    }
    public struct Point11123
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 29;
    }

    public static class Grid11123
    {
        public static Point11123 Make(int left, int right)
        {
            return new Point11123 { X = left, Y = right };
        }
    }
    public delegate void Notifier11124(int value);

    public sealed class Source11124
    {
        public event Notifier11124 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point11125
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 43;
    }

    public static class Grid11125
    {
        public static Point11125 Make(int left, int right)
        {
            return new Point11125 { X = left, Y = right };
        }
    }
}
