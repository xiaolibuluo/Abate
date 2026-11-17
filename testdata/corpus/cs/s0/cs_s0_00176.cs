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

    public struct Point1680
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 55;
    }

    public static class Grid1680
    {
        public static Point1680 Make(int left, int right)
        {
            return new Point1680 { X = left, Y = right };
        }
    }
    public delegate void Notifier1681(int value);

    public sealed class Source1681
    {
        public event Notifier1681 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler1682
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope1682(string Label, int Payload);

    public sealed class Handler1682 : IHandler1682
    {
        public void Handle(int value) { }
    }
}
