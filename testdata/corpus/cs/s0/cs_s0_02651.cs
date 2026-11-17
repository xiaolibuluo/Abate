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

    public struct Point24800
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 41;
    }

    public static class Grid24800
    {
        public static Point24800 Make(int left, int right)
        {
            return new Point24800 { X = left, Y = right };
        }
    }
    public interface IHandler24801
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope24801(string Label, int Payload);

    public sealed class Handler24801 : IHandler24801
    {
        public void Handle(int value) { }
    }
    public struct Point24802
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 55;
    }

    public static class Grid24802
    {
        public static Point24802 Make(int left, int right)
        {
            return new Point24802 { X = left, Y = right };
        }
    }
    public static int Scale24803<T>(T value) where T : struct
    {
        int number = 62;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 17; }
        return number;
    }
    public interface IHandler24804
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope24804(string Label, int Payload);

    public sealed class Handler24804 : IHandler24804
    {
        public void Handle(int value) { }
    }
    public struct Point24805
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 76;
    }

    public static class Grid24805
    {
        public static Point24805 Make(int left, int right)
        {
            return new Point24805 { X = left, Y = right };
        }
    }
}
