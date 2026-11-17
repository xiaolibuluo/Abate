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

    public delegate void Notifier9000(int value);

    public sealed class Source9000
    {
        public event Notifier9000 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier9001(int value);

    public sealed class Source9001
    {
        public event Notifier9001 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point9002
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 79;
    }

    public static class Grid9002
    {
        public static Point9002 Make(int left, int right)
        {
            return new Point9002 { X = left, Y = right };
        }
    }
    public interface IHandler9003
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope9003(string Label, int Payload);

    public sealed class Handler9003 : IHandler9003
    {
        public void Handle(int value) { }
    }
}
