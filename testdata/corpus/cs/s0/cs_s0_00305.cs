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

    public interface IHandler2880
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope2880(string Label, int Payload);

    public sealed class Handler2880 : IHandler2880
    {
        public void Handle(int value) { }
    }
    public interface IHandler2881
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope2881(string Label, int Payload);

    public sealed class Handler2881 : IHandler2881
    {
        public void Handle(int value) { }
    }
    public static int Scale2882<T>(T value) where T : struct
    {
        int number = 77;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 34; }
        return number;
    }
    public static int Scale2883<T>(T value) where T : struct
    {
        int number = 84;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 45; }
        return number;
    }
    public struct Point2884
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 91;
    }

    public static class Grid2884
    {
        public static Point2884 Make(int left, int right)
        {
            return new Point2884 { X = left, Y = right };
        }
    }
    public static int Scale2885<T>(T value) where T : struct
    {
        int number = 98;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 67; }
        return number;
    }
    public struct Point2886
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 8;
    }

    public static class Grid2886
    {
        public static Point2886 Make(int left, int right)
        {
            return new Point2886 { X = left, Y = right };
        }
    }
}
