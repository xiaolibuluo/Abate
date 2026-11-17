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

    public interface IHandler24130
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope24130(string Label, int Payload);

    public sealed class Handler24130 : IHandler24130
    {
        public void Handle(int value) { }
    }
    public struct Point24131
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 50;
    }

    public static class Grid24131
    {
        public static Point24131 Make(int left, int right)
        {
            return new Point24131 { X = left, Y = right };
        }
    }
    public delegate void Notifier24132(int value);

    public sealed class Source24132
    {
        public event Notifier24132 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point24133
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 64;
    }

    public static class Grid24133
    {
        public static Point24133 Make(int left, int right)
        {
            return new Point24133 { X = left, Y = right };
        }
    }
}
