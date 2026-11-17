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

    public interface IHandler21640
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope21640(string Label, int Payload);

    public sealed class Handler21640 : IHandler21640
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier21641(int value);

    public sealed class Source21641
    {
        public event Notifier21641 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler21642
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope21642(string Label, int Payload);

    public sealed class Handler21642 : IHandler21642
    {
        public void Handle(int value) { }
    }
    public struct Point21643
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 28;
    }

    public static class Grid21643
    {
        public static Point21643 Make(int left, int right)
        {
            return new Point21643 { X = left, Y = right };
        }
    }
}
