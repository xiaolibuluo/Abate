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

    public interface IHandler13890
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope13890(string Label, int Payload);

    public sealed class Handler13890 : IHandler13890
    {
        public void Handle(int value) { }
    }
    public struct Point13891
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 27;
    }

    public static class Grid13891
    {
        public static Point13891 Make(int left, int right)
        {
            return new Point13891 { X = left, Y = right };
        }
    }
    public static int Scale13892<T>(T value) where T : struct
    {
        int number = 34;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 55; }
        return number;
    }
    public interface IHandler13893
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope13893(string Label, int Payload);

    public sealed class Handler13893 : IHandler13893
    {
        public void Handle(int value) { }
    }
}
