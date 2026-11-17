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

    public struct Point16670
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 45;
    }

    public static class Grid16670
    {
        public static Point16670 Make(int left, int right)
        {
            return new Point16670 { X = left, Y = right };
        }
    }
    public interface IHandler16671
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope16671(string Label, int Payload);

    public sealed class Handler16671 : IHandler16671
    {
        public void Handle(int value) { }
    }
}
