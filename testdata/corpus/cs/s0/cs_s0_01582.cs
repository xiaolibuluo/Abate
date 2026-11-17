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

    public interface IHandler14800
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope14800(string Label, int Payload);

    public sealed class Handler14800 : IHandler14800
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier14801(int value);

    public sealed class Source14801
    {
        public event Notifier14801 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point14802
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 53;
    }

    public static class Grid14802
    {
        public static Point14802 Make(int left, int right)
        {
            return new Point14802 { X = left, Y = right };
        }
    }
}
