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

    public interface IHandler15250
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15250(string Label, int Payload);

    public sealed class Handler15250 : IHandler15250
    {
        public void Handle(int value) { }
    }
    public struct Point15251
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 49;
    }

    public static class Grid15251
    {
        public static Point15251 Make(int left, int right)
        {
            return new Point15251 { X = left, Y = right };
        }
    }
}
