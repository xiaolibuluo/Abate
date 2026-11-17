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

    public interface IHandler18590
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope18590(string Label, int Payload);

    public sealed class Handler18590 : IHandler18590
    {
        public void Handle(int value) { }
    }
    public struct Point18591
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 26;
    }

    public static class Grid18591
    {
        public static Point18591 Make(int left, int right)
        {
            return new Point18591 { X = left, Y = right };
        }
    }
}
