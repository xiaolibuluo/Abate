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

    public struct Point10850
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 45;
    }

    public static class Grid10850
    {
        public static Point10850 Make(int left, int right)
        {
            return new Point10850 { X = left, Y = right };
        }
    }
    public interface IHandler10851
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope10851(string Label, int Payload);

    public sealed class Handler10851 : IHandler10851
    {
        public void Handle(int value) { }
    }
}
