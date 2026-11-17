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

    public interface IHandler20670
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope20670(string Label, int Payload);

    public sealed class Handler20670 : IHandler20670
    {
        public void Handle(int value) { }
    }
    public static int Scale20671<T>(T value) where T : struct
    {
        int number = 14;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 32; }
        return number;
    }
    public struct Point20672
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 21;
    }

    public static class Grid20672
    {
        public static Point20672 Make(int left, int right)
        {
            return new Point20672 { X = left, Y = right };
        }
    }
}
