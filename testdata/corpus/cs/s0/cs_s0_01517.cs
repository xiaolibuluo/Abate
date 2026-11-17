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

    public interface IHandler14210
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope14210(string Label, int Payload);

    public sealed class Handler14210 : IHandler14210
    {
        public void Handle(int value) { }
    }
    public struct Point14211
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 55;
    }

    public static class Grid14211
    {
        public static Point14211 Make(int left, int right)
        {
            return new Point14211 { X = left, Y = right };
        }
    }
}
