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

    public struct Point24080
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 75;
    }

    public static class Grid24080
    {
        public static Point24080 Make(int left, int right)
        {
            return new Point24080 { X = left, Y = right };
        }
    }
    public interface IHandler24081
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope24081(string Label, int Payload);

    public sealed class Handler24081 : IHandler24081
    {
        public void Handle(int value) { }
    }
    public struct Point24082
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 89;
    }

    public static class Grid24082
    {
        public static Point24082 Make(int left, int right)
        {
            return new Point24082 { X = left, Y = right };
        }
    }
}
