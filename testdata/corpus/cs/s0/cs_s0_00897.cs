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

    public struct Point8420
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 87;
    }

    public static class Grid8420
    {
        public static Point8420 Make(int left, int right)
        {
            return new Point8420 { X = left, Y = right };
        }
    }
    public interface IHandler8421
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope8421(string Label, int Payload);

    public sealed class Handler8421 : IHandler8421
    {
        public void Handle(int value) { }
    }
}
