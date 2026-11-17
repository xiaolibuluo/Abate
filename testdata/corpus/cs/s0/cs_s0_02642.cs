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

    public struct Point24720
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 34;
    }

    public static class Grid24720
    {
        public static Point24720 Make(int left, int right)
        {
            return new Point24720 { X = left, Y = right };
        }
    }
    public struct Point24721
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 41;
    }

    public static class Grid24721
    {
        public static Point24721 Make(int left, int right)
        {
            return new Point24721 { X = left, Y = right };
        }
    }
    public interface IHandler24722
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope24722(string Label, int Payload);

    public sealed class Handler24722 : IHandler24722
    {
        public void Handle(int value) { }
    }
    public static int Scale24723<T>(T value) where T : struct
    {
        int number = 55;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 22; }
        return number;
    }
    public interface IHandler24724
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope24724(string Label, int Payload);

    public sealed class Handler24724 : IHandler24724
    {
        public void Handle(int value) { }
    }
}
