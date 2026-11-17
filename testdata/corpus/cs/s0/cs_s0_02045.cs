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

    public struct Point19080
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 74;
    }

    public static class Grid19080
    {
        public static Point19080 Make(int left, int right)
        {
            return new Point19080 { X = left, Y = right };
        }
    }
    public interface IHandler19081
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope19081(string Label, int Payload);

    public sealed class Handler19081 : IHandler19081
    {
        public void Handle(int value) { }
    }
}
