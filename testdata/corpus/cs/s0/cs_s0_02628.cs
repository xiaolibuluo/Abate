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

    public interface IHandler24580
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope24580(string Label, int Payload);

    public sealed class Handler24580 : IHandler24580
    {
        public void Handle(int value) { }
    }
    public struct Point24581
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 53;
    }

    public static class Grid24581
    {
        public static Point24581 Make(int left, int right)
        {
            return new Point24581 { X = left, Y = right };
        }
    }
    public struct Point24582
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 60;
    }

    public static class Grid24582
    {
        public static Point24582 Make(int left, int right)
        {
            return new Point24582 { X = left, Y = right };
        }
    }
}
