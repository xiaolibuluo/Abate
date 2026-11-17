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

    public struct Point23620
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 59;
    }

    public static class Grid23620
    {
        public static Point23620 Make(int left, int right)
        {
            return new Point23620 { X = left, Y = right };
        }
    }
    public interface IHandler23621
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope23621(string Label, int Payload);

    public sealed class Handler23621 : IHandler23621
    {
        public void Handle(int value) { }
    }
    public interface IHandler23622
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope23622(string Label, int Payload);

    public sealed class Handler23622 : IHandler23622
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier23623(int value);

    public sealed class Source23623
    {
        public event Notifier23623 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
