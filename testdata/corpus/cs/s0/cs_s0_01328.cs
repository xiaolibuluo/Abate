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

    public delegate void Notifier12440(int value);

    public sealed class Source12440
    {
        public event Notifier12440 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point12441
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 82;
    }

    public static class Grid12441
    {
        public static Point12441 Make(int left, int right)
        {
            return new Point12441 { X = left, Y = right };
        }
    }
    public interface IHandler12442
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope12442(string Label, int Payload);

    public sealed class Handler12442 : IHandler12442
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier12443(int value);

    public sealed class Source12443
    {
        public event Notifier12443 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
