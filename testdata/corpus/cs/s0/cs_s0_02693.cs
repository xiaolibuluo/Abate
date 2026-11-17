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

    public static int Scale25180<T>(T value) where T : struct
    {
        int number = 50;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 67; }
        return number;
    }
    public struct Point25181
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 57;
    }

    public static class Grid25181
    {
        public static Point25181 Make(int left, int right)
        {
            return new Point25181 { X = left, Y = right };
        }
    }
    public struct Point25182
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 64;
    }

    public static class Grid25182
    {
        public static Point25182 Make(int left, int right)
        {
            return new Point25182 { X = left, Y = right };
        }
    }
    public interface IHandler25183
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope25183(string Label, int Payload);

    public sealed class Handler25183 : IHandler25183
    {
        public void Handle(int value) { }
    }
    public interface IHandler25184
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope25184(string Label, int Payload);

    public sealed class Handler25184 : IHandler25184
    {
        public void Handle(int value) { }
    }
}
