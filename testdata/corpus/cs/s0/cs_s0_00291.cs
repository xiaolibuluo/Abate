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

    public delegate void Notifier2740(int value);

    public sealed class Source2740
    {
        public event Notifier2740 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler2741
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope2741(string Label, int Payload);

    public sealed class Handler2741 : IHandler2741
    {
        public void Handle(int value) { }
    }
    public struct Point2742
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 89;
    }

    public static class Grid2742
    {
        public static Point2742 Make(int left, int right)
        {
            return new Point2742 { X = left, Y = right };
        }
    }
    public interface IHandler2743
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope2743(string Label, int Payload);

    public sealed class Handler2743 : IHandler2743
    {
        public void Handle(int value) { }
    }
    public static int Scale2744<T>(T value) where T : struct
    {
        int number = 6;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 19; }
        return number;
    }
}
