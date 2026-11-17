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

    public delegate void Notifier5280(int value);

    public sealed class Source5280
    {
        public event Notifier5280 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier5281(int value);

    public sealed class Source5281
    {
        public event Notifier5281 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler5282
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5282(string Label, int Payload);

    public sealed class Handler5282 : IHandler5282
    {
        public void Handle(int value) { }
    }
    public interface IHandler5283
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5283(string Label, int Payload);

    public sealed class Handler5283 : IHandler5283
    {
        public void Handle(int value) { }
    }
    public interface IHandler5284
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5284(string Label, int Payload);

    public sealed class Handler5284 : IHandler5284
    {
        public void Handle(int value) { }
    }
    public struct Point5285
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 17;
    }

    public static class Grid5285
    {
        public static Point5285 Make(int left, int right)
        {
            return new Point5285 { X = left, Y = right };
        }
    }
}
