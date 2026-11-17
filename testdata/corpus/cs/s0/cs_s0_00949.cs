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

    public interface IHandler8890
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope8890(string Label, int Payload);

    public sealed class Handler8890 : IHandler8890
    {
        public void Handle(int value) { }
    }
    public struct Point8891
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 26;
    }

    public static class Grid8891
    {
        public static Point8891 Make(int left, int right)
        {
            return new Point8891 { X = left, Y = right };
        }
    }
}
