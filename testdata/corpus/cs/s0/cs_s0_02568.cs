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

    public struct Point24010
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 81;
    }

    public static class Grid24010
    {
        public static Point24010 Make(int left, int right)
        {
            return new Point24010 { X = left, Y = right };
        }
    }
    public interface IHandler24011
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope24011(string Label, int Payload);

    public sealed class Handler24011 : IHandler24011
    {
        public void Handle(int value) { }
    }
}
