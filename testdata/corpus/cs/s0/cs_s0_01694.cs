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

    public struct Point15830
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 20;
    }

    public static class Grid15830
    {
        public static Point15830 Make(int left, int right)
        {
            return new Point15830 { X = left, Y = right };
        }
    }
    public struct Point15831
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 27;
    }

    public static class Grid15831
    {
        public static Point15831 Make(int left, int right)
        {
            return new Point15831 { X = left, Y = right };
        }
    }
    public interface IHandler15832
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15832(string Label, int Payload);

    public sealed class Handler15832 : IHandler15832
    {
        public void Handle(int value) { }
    }
}
