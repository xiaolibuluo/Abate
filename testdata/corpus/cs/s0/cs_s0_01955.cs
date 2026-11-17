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

    public delegate void Notifier18230(int value);

    public sealed class Source18230
    {
        public event Notifier18230 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler18231
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope18231(string Label, int Payload);

    public sealed class Handler18231 : IHandler18231
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier18232(int value);

    public sealed class Source18232
    {
        public event Notifier18232 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point18233
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 57;
    }

    public static class Grid18233
    {
        public static Point18233 Make(int left, int right)
        {
            return new Point18233 { X = left, Y = right };
        }
    }
}
