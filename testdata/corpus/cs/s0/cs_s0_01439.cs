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

    public struct Point13480
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 69;
    }

    public static class Grid13480
    {
        public static Point13480 Make(int left, int right)
        {
            return new Point13480 { X = left, Y = right };
        }
    }
    public delegate void Notifier13481(int value);

    public sealed class Source13481
    {
        public event Notifier13481 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point13482
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 83;
    }

    public static class Grid13482
    {
        public static Point13482 Make(int left, int right)
        {
            return new Point13482 { X = left, Y = right };
        }
    }
    public static int Scale13483<T>(T value) where T : struct
    {
        int number = 90;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 23; }
        return number;
    }
    public interface IHandler13484
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope13484(string Label, int Payload);

    public sealed class Handler13484 : IHandler13484
    {
        public void Handle(int value) { }
    }
}
