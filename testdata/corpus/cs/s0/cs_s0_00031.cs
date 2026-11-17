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

    public delegate void Notifier290(int value);

    public sealed class Source290
    {
        public event Notifier290 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point291
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 98;
    }

    public static class Grid291
    {
        public static Point291 Make(int left, int right)
        {
            return new Point291 { X = left, Y = right };
        }
    }
    public interface IHandler292
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope292(string Label, int Payload);

    public sealed class Handler292 : IHandler292
    {
        public void Handle(int value) { }
    }
    public interface IHandler293
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope293(string Label, int Payload);

    public sealed class Handler293 : IHandler293
    {
        public void Handle(int value) { }
    }
}
