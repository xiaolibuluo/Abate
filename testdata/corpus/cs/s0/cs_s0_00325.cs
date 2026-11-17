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

    public struct Point3070
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 19;
    }

    public static class Grid3070
    {
        public static Point3070 Make(int left, int right)
        {
            return new Point3070 { X = left, Y = right };
        }
    }
    public interface IHandler3071
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope3071(string Label, int Payload);

    public sealed class Handler3071 : IHandler3071
    {
        public void Handle(int value) { }
    }
}
