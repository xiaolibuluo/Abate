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

    public delegate void Notifier26440(int value);

    public sealed class Source26440
    {
        public event Notifier26440 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler26441
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope26441(string Label, int Payload);

    public sealed class Handler26441 : IHandler26441
    {
        public void Handle(int value) { }
    }
    public struct Point26442
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 53;
    }

    public static class Grid26442
    {
        public static Point26442 Make(int left, int right)
        {
            return new Point26442 { X = left, Y = right };
        }
    }
}
