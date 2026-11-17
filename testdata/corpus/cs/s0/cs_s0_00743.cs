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

    public interface IHandler6970
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope6970(string Label, int Payload);

    public sealed class Handler6970 : IHandler6970
    {
        public void Handle(int value) { }
    }
    public struct Point6971
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 52;
    }

    public static class Grid6971
    {
        public static Point6971 Make(int left, int right)
        {
            return new Point6971 { X = left, Y = right };
        }
    }
}
