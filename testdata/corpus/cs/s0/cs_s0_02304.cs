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

    public interface IHandler21500
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope21500(string Label, int Payload);

    public sealed class Handler21500 : IHandler21500
    {
        public void Handle(int value) { }
    }
    public struct Point21501
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 26;
    }

    public static class Grid21501
    {
        public static Point21501 Make(int left, int right)
        {
            return new Point21501 { X = left, Y = right };
        }
    }
}
