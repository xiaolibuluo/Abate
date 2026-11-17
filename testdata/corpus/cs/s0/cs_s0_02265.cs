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

    public interface IHandler21150
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope21150(string Label, int Payload);

    public sealed class Handler21150 : IHandler21150
    {
        public void Handle(int value) { }
    }
    public struct Point21151
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 56;
    }

    public static class Grid21151
    {
        public static Point21151 Make(int left, int right)
        {
            return new Point21151 { X = left, Y = right };
        }
    }
    public interface IHandler21152
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope21152(string Label, int Payload);

    public sealed class Handler21152 : IHandler21152
    {
        public void Handle(int value) { }
    }
}
