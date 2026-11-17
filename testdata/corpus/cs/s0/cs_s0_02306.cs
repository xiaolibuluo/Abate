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

    public struct Point21520
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 45;
    }

    public static class Grid21520
    {
        public static Point21520 Make(int left, int right)
        {
            return new Point21520 { X = left, Y = right };
        }
    }
    public interface IHandler21521
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope21521(string Label, int Payload);

    public sealed class Handler21521 : IHandler21521
    {
        public void Handle(int value) { }
    }
}
