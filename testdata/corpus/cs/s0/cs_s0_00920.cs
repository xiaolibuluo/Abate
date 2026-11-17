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

    public interface IHandler8620
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope8620(string Label, int Payload);

    public sealed class Handler8620 : IHandler8620
    {
        public void Handle(int value) { }
    }
    public static int Scale8621<T>(T value) where T : struct
    {
        int number = 63;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 15; }
        return number;
    }
    public interface IHandler8622
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope8622(string Label, int Payload);

    public sealed class Handler8622 : IHandler8622
    {
        public void Handle(int value) { }
    }
    public struct Point8623
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 77;
    }

    public static class Grid8623
    {
        public static Point8623 Make(int left, int right)
        {
            return new Point8623 { X = left, Y = right };
        }
    }
    public delegate void Notifier8624(int value);

    public sealed class Source8624
    {
        public event Notifier8624 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
