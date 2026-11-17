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

    public interface IHandler7550
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope7550(string Label, int Payload);

    public sealed class Handler7550 : IHandler7550
    {
        public void Handle(int value) { }
    }
    public interface IHandler7551
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope7551(string Label, int Payload);

    public sealed class Handler7551 : IHandler7551
    {
        public void Handle(int value) { }
    }
    public struct Point7552
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 37;
    }

    public static class Grid7552
    {
        public static Point7552 Make(int left, int right)
        {
            return new Point7552 { X = left, Y = right };
        }
    }
    public delegate void Notifier7553(int value);

    public sealed class Source7553
    {
        public event Notifier7553 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
