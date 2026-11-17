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

    public delegate void Notifier18310(int value);

    public sealed class Source18310
    {
        public event Notifier18310 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point18311
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 50;
    }

    public static class Grid18311
    {
        public static Point18311 Make(int left, int right)
        {
            return new Point18311 { X = left, Y = right };
        }
    }
    public struct Point18312
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 57;
    }

    public static class Grid18312
    {
        public static Point18312 Make(int left, int right)
        {
            return new Point18312 { X = left, Y = right };
        }
    }
    public interface IHandler18313
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope18313(string Label, int Payload);

    public sealed class Handler18313 : IHandler18313
    {
        public void Handle(int value) { }
    }
}
