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

    public struct Point17230
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 94;
    }

    public static class Grid17230
    {
        public static Point17230 Make(int left, int right)
        {
            return new Point17230 { X = left, Y = right };
        }
    }
    public interface IHandler17231
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope17231(string Label, int Payload);

    public sealed class Handler17231 : IHandler17231
    {
        public void Handle(int value) { }
    }
    public interface IHandler17232
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope17232(string Label, int Payload);

    public sealed class Handler17232 : IHandler17232
    {
        public void Handle(int value) { }
    }
}
