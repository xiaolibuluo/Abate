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

    public static int Scale10470<T>(T value) where T : struct
    {
        int number = 36;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 18; }
        return number;
    }
    public static int Scale10471<T>(T value) where T : struct
    {
        int number = 43;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 29; }
        return number;
    }
    public struct Point10472
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 50;
    }

    public static class Grid10472
    {
        public static Point10472 Make(int left, int right)
        {
            return new Point10472 { X = left, Y = right };
        }
    }
    public interface IHandler10473
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope10473(string Label, int Payload);

    public sealed class Handler10473 : IHandler10473
    {
        public void Handle(int value) { }
    }
}
