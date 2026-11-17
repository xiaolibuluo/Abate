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

    public struct Point4990
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 90;
    }

    public static class Grid4990
    {
        public static Point4990 Make(int left, int right)
        {
            return new Point4990 { X = left, Y = right };
        }
    }
    public interface IHandler4991
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4991(string Label, int Payload);

    public sealed class Handler4991 : IHandler4991
    {
        public void Handle(int value) { }
    }
}
