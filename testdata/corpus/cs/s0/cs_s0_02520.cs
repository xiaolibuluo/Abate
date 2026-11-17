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

    public struct Point23540
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 52;
    }

    public static class Grid23540
    {
        public static Point23540 Make(int left, int right)
        {
            return new Point23540 { X = left, Y = right };
        }
    }
    public interface IHandler23541
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope23541(string Label, int Payload);

    public sealed class Handler23541 : IHandler23541
    {
        public void Handle(int value) { }
    }
    public struct Point23542
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 66;
    }

    public static class Grid23542
    {
        public static Point23542 Make(int left, int right)
        {
            return new Point23542 { X = left, Y = right };
        }
    }
}
