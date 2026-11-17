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

    public struct Point4610
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 81;
    }

    public static class Grid4610
    {
        public static Point4610 Make(int left, int right)
        {
            return new Point4610 { X = left, Y = right };
        }
    }
    public delegate void Notifier4611(int value);

    public sealed class Source4611
    {
        public event Notifier4611 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier4612(int value);

    public sealed class Source4612
    {
        public event Notifier4612 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler4613
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4613(string Label, int Payload);

    public sealed class Handler4613 : IHandler4613
    {
        public void Handle(int value) { }
    }
}
